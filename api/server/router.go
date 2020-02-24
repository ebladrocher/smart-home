package server

import (
	"github.com/ebladrocher/smarthome/api/server/handlers"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/gorilla/mux"
)

// setHandlers ...
func SetHandlers(router *mux.Router, uc store.UseCase)  {

	h := handlers.NewHandlers(uc)

	// Index
	router.HandleFunc("/", h.Index.HandleIndex()).Methods("GET")

	// User
	router.HandleFunc("/auth/sign-up", h.User.SignUp).Methods("POST")

}


