package gapi

import (
	"github.com/axilock/axilock-backend/internal/auth"
	"github.com/axilock/axilock-backend/internal/db/dynamo"
	"github.com/axilock/axilock-backend/internal/db/s3store"
	"github.com/axilock/axilock-backend/internal/service"
	"github.com/axilock/axilock-backend/pkg/util"
	"github.com/axilock/axilock-backend/pkg/workers"
)

type Server struct {
	TokenMaker      auth.Maker
	DynamoStore     dynamo.NoSQLStoreInterface
	S3store         s3store.S3StoreInterface
	Config          util.Config
	Services        *service.Services
	TaskDistributer workers.TaskDistributerInterface
}

type ServerConfig struct {
	TokenMaker      auth.Maker
	DynamoStore     dynamo.NoSQLStoreInterface
	S3store         s3store.S3StoreInterface
	Config          util.Config
	Services        *service.Services
	TaskDistributer workers.TaskDistributerInterface
}

func NewServer(cfg ServerConfig) (*Server, error) {
	server := &Server{
		TokenMaker:      cfg.TokenMaker,
		Config:          cfg.Config,
		Services:        cfg.Services,
		DynamoStore:     cfg.DynamoStore,
		S3store:         cfg.S3store,
		TaskDistributer: cfg.TaskDistributer,
	}
	return server, nil
}
