package api

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleIndex(t *testing.T)  {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	cfg := config.AppConfig{
		ServerHost:"0.0.0.0",
		ServerPort: 3000,
		Mode: config.DebugMode,
	}

	logConfig := server.InitLogger(&cfg)

	store := teststore.NewStore()

	srvConf := server.NewServerConfig(&cfg)

	srv, err := server.NewServer(srvConf, store, logConfig)
	if err != nil {
		panic(err.Error())
	}

	srv.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
