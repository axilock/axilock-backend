package api

// func newTestServer(t *testing.T, store db.Store) *Server {
// 	config := util.Config{
// 		SymetricKey:         util.RandString(32),
// 		AccessTokenDuration: time.Minute,
// 		RunningEnv:          "test",
// 	}
// 	tokenMaker, err := auth.NewPasetoMaker(config.SymetricKey)
// 	if err != nil {
// 		log.Fatal().Msg("cannot init tokenmaker")
// 	}
// 	services := service.NewServiceInit(store, config, redis.RedisStore{})
// 	srvcfg := ServerConfig{
// 		TokenMaker: tokenMaker,
// 		Config:     config,
// 		Services:   services,
// 	}
// 	server, err := NewServer(srvcfg)
// 	require.NoError(t, err)
// 	return server
// }
