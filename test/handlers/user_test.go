package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/ebladrocher/smarthome/adaptors"
	"github.com/ebladrocher/smarthome/api/server"
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

	type signInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	signUpBody := &signInput{
		Email: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Email, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
