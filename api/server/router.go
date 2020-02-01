package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	router *mux.Router
}

func New() *Router{
	var s = Router{router:mux.NewRouter()}
	return &s
}

func (s *Router) configureRouter()  {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *Router) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("☄ HTTP status code returned!"))
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

/*func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	// process the request...
}
*/