package handlers

import (
	"encoding/json"
	"github.com/ebladrocher/smarthome/system/store"
	"net/http"
)

type HandlerUser struct {
	useCase store.UseCase
}

func NewHandlerUser(useCase store.UseCase) *HandlerUser {
	return &HandlerUser{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *HandlerUser) SignUp(w http.ResponseWriter, r *http.Request) {
	inp := new(signInput)
	json.NewDecoder(r.Body).Decode(inp)
	//json.NewEncoder(w).Encode(p)

	if err := h.useCase.SignUp(inp.Username, inp.Password); err != nil {
		panic(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
