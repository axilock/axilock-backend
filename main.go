package main

import (
	"context"
	"net"
	"os"

	"github.com/alecthomas/kong"
	"github.com/axilock/axilock-backend/internal/api"
	"github.com/axilock/axilock-backend/internal/auth"
	"github.com/axilock/axilock-backend/internal/cli"
	"github.com/axilock/axilock-backend/internal/db/dynamo"
	"github.com/axilock/axilock-backend/internal/db/redis"
	"github.com/axilock/axilock-backend/internal/db/s3store"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/axilock/axilock-backend/internal/gapi"
	"github.com/axilock/axilock-backend/internal/gapi/clientrpc"
	"github.com/axilock/axilock-backend/internal/gapi/runnerpc"
	"github.com/axilock/axilock-backend/internal/service"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/axilock/axilock-backend/pkg/workers"
	clientpb "github.com/axilock/axilock-protos/client"
	runnerpb "github.com/axilock/axilock-protos/runner"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig()
	if config.RunningEnv != "prod" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}
	defer conn.Close()

	store := db.NewStore(conn)
	dynamoStore := dynamo.NewDynamoDBStore()
	redis := redis.NewRedisClient(config.RedisAddr, config.RedisPassword)
	services := service.NewServiceInit(store, config, redis)
	tokenMaker := auth.NewNativeAuth(redis)
	if err != nil {
		log.Fatal().Msg("cannot init tokenmaker")
	}
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init githubclient")
	}
	s3store := s3store.NewS3Clinet()
	redisOptn := asynq.RedisClientOpt{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
	}
	taskDistributer := workers.NewRedisTaskDistributer(redisOptn, *services)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init superset")
	}
	grpccfg := gapi.ServerConfig{
		DynamoStore:     dynamoStore,
		Services:        services,
		TokenMaker:      tokenMaker,
		Config:          config,
		S3store:         s3store,
		TaskDistributer: taskDistributer,
	}

	gincfg := api.ServerConfig{
		TokenMaker:      tokenMaker,
		Config:          config,
		Services:        services,
		TaskDistributer: taskDistributer,
	}

	ctx := kong.Parse(&cli.Cli, kong.BindTo(store, (*db.Store)(nil)))
	if ctx.Command() != "serve" {
		err = ctx.Run()
		if err != nil {
			log.Error().Err(err).Msg("cannot run command")
			os.Exit(1)
		}
		return
	}

	go workers.InitProcessor(redisOptn, *services)
	go runGinServer(gincfg)
	runGRPCServer(grpccfg)
}

func runGRPCServer(cfg gapi.ServerConfig) {
	server, err := gapi.NewServer(cfg)
	if err != nil {
		log.Error().Err(err).Msg("cannot create server")
		return
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GRPCLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	// ------------------------------------
	// Creating cliinet RPC group and register

	clientGrp := &clientrpc.ClientService{Server: server}

	clientpb.RegisterHealthServiceServer(grpcServer, clientGrp)
	clientpb.RegisterSessionServiceServer(grpcServer, clientGrp)
	clientpb.RegisterMetadataServiceServer(grpcServer, clientGrp)
	clientpb.RegisterRegexServiceServer(grpcServer, clientGrp)
	clientpb.RegisterCommitDataServiceServer(grpcServer, clientGrp)
	clientpb.RegisterAlertServiceServer(grpcServer, clientGrp)

	// ------------------------------------
	// Creating runner  RPC group and register
	runnerGrp := &runnerpc.RunnerService{Server: server}

	runnerpb.RegisterScanServiceServer(grpcServer, runnerGrp)

	// ------------------------------------
	if cfg.Config.RunningEnv != "prod" {
		reflection.Register(grpcServer)
	}

	listener, err := net.Listen("tcp", cfg.Config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot start gRPC server")
		return
	}
	log.Info().Msgf("staring gRPC Server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("cannot start gRPC server")
	}
}

// func runGatewayServer(cfg gapi.ServerConfig) {
// 	server, err := gapi.NewServer(cfg)
// 	if err != nil {
// 		log.Fatal().Msg("cannot create server")
// 	}
// 	jsonOptn := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		MarshalOptions: protojson.MarshalOptions{
// 			UseProtoNames: true,
// 		},
// 		UnmarshalOptions: protojson.UnmarshalOptions{
// 			DiscardUnknown: true,
// 		},
// 	})
// 	grpcMux := runtime.NewServeMux(jsonOptn)

// 	ctx, cancel := context.WithCancel(context.Background())

// 	defer cancel()

// 	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, server)
// 	if err != nil {
// 		log.Fatal().Msg("cannot register handler server") //nolint:gocritic
// 	}
// 	mux := http.NewServeMux()

// 	mux.Handle("/", grpcMux)

// 	listener, err := net.Listen("tcp", cfg.Config.GatewayServerAddr)
// 	if err != nil {
// 		log.Fatal().Msg("cannot start gRPC server")
// 	}

// 	muxserver := http.Server{
// 		Handler:      mux,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  30 * time.Second,
// 	}

// 	log.Info().Msgf("staring HTTP Gateway Server at %s", listener.Addr().String())
// 	err = muxserver.Serve(listener)
// 	if err != nil {
// 		log.Fatal().AnErr("error", err).Msg("cannot start gRPC server")
// 	}
// }

func runGinServer(cfg api.ServerConfig) {
	server, err := api.NewServer(cfg)
	if err != nil {
		log.Error().Err(err).Msg("cannot create server")
		return
	}
	err = server.Start(cfg.Config.HTTPServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot start gin server")
	}
}
