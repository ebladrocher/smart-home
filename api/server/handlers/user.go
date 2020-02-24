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

type SignInput struct {
	Email string `json:"username"`
	Password string `json:"password"`
}

func (h *HandlerUser) SignUp(w http.ResponseWriter, r *http.Request) {
	inp := new(SignInput)
	json.NewDecoder(r.Body).Decode(inp)
	//json.NewEncoder(w).Encode(p)

	if err := h.useCase.SignUp(inp.Email, inp.Password); err != nil {
		panic(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
