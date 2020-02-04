package server

import (
	"fmt"
	"net/http"
	"smarthome/system/store"
)

type Server struct {
	config *ServerConfig
	server *http.Server
	router *Router
	store *store.Store
}

func (s *Server) Start() error{
	s.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler: s.router.router,
	}

	s.router.configureRouter()

	if err := s.store.ConfigureStore(); err != nil {
		return err
	}

	return s.server.ListenAndServe()
}

func NewServer(
	/*cfg *ServerConfig,
	*db *store.Store,*/
	) (newServer *Server) {
	cfg := NewServerConfig()
	newServer = &Server{
		config: cfg,
		router: New(),
		store: store.Init(),
	}

	return
}

