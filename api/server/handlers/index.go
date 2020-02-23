package handlers

import (
	"net/http"
)

// ControllerIndex ...
type HandlerIndex struct {

}

// NewControllerIndex ...
func NewHandlerIndex() *HandlerIndex {
	return &HandlerIndex{}
}

// HandleIndex ...
func (i HandlerIndex ) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("☄ HTTP status code returned!"))
	}
}