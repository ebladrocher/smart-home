package handlers

import (
	"encoding/json"
	"errors"
	"github.com/ebladrocher/smarthome/system/store"
	"net/http"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

// HandlerUser ...
type HandlerUser struct {
	useCase store.UseCase
}

// NewHandlerUser ...
func NewHandlerUser(useCase store.UseCase) *HandlerUser {
	return &HandlerUser{
		useCase: useCase,
	}
}

// SignUp ...
func (h *HandlerUser) SignUp() http.HandlerFunc {
	type signInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		inp := new(signInput)
		if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
			newError(w,r, http.StatusBadRequest, err)
		}

		if err := h.useCase.SignUp(inp.Email, inp.Password); err != nil {
			newError(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		Respond(w, r, http.StatusCreated, inp)
	}
}

// SignIn ...
func (h *HandlerUser) SignIn() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		inp := new(request)
		if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
			newError(w,r, http.StatusBadRequest, err)
		}

		_ ,err := h.useCase.SignIn(inp.Email, inp.Password)
		if err != nil {
			newError(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		Respond(w, r, http.StatusOK, nil)
	}
}

