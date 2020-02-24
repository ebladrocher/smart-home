package api

import (
	"bytes"
	"encoding/json"
	"github.com/ebladrocher/smarthome/adaptors"
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/api/server/handlers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	r := mux.NewRouter()
	uc := new(adaptors.AuthUseCaseMock)

	server.SetHandlers(r, uc)

	signUpBody := &handlers.SignInput{
		Email: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Email, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

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

