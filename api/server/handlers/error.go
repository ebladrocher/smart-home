package handlers

import "net/http"

func newError(w http.ResponseWriter, r *http.Request, code int, err error)  {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}
