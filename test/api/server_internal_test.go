package api

//func TestServer_HandleIndex(t *testing.T)  {
//	rec := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodGet, "/", nil)
//
//	cfg := config.AppConfig{
//		ServerHost:"0.0.0.0",
//		ServerPort: 3000,
//		Mode: config.DebugMode,
//	}
//
//	logConfig := server.InitLogger(&cfg)
//
//
//	srvConf := server.NewServerConfig(&cfg)
//
//	srv, err := server.NewServer(srvConf, logConfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	srv.ServeHTTP(rec, req)
//	assert.Equal(t, rec.Code, http.StatusOK)
//}

