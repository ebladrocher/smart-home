package controllers

import (
	"net/http"
)

// ControllerIndex ...
type ControllerIndex struct {

}

// NewControllerIndex ...
func NewControllerIndex() *ControllerIndex {
	return &ControllerIndex{}
}

// HandleIndex ...
func (i ControllerIndex ) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("☄ HTTP status code returned!"))
	}
}