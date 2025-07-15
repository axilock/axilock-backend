package gapi

// func newTestServer(t *testing.T, store db.Store) *Server {
// 	config := util.Config{
// 		SymetricKey:         util.RandString(32),
// 		AccessTokenDuration: time.Minute,
// 		RunningEnv:          "dev",
// 	}
// 	tokenMaker, _ := auth.NewPasetoMaker(config.SymetricKey)
// 	services := service.NewServiceInit(store, config, redis.RedisStore{})
// 	cfg := ServerConfig{
// 		TokenMaker: tokenMaker,
// 		Config:     config,
// 		Services:   services,
// 	}
// 	server, err := NewServer(cfg)
// 	require.NoError(t, err)
// 	return server
// }
