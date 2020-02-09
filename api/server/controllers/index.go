package controllers

import (
	"net/http"
)

type ControllerIndex struct {

}

func NewControllerIndex() *ControllerIndex {
	return &ControllerIndex{}
}

func (i ControllerIndex ) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("☄ HTTP status code returned!"))
	}
}