package server

import (
	"fmt"
	"net/http"
	"github.com/ebladrocher/smarthome/system/store"
)



type Server struct {
	config *ServerConfig
	server *http.Server
	router *Router
	store *store.Store
	logger *ServerLogger
}

func (s *Server) Start() error{
	s.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler: s.router.router,
	}

	s.router.ConfigureRouter()

	if err := s.store.ConfigureStore(); err != nil {
		return err
	}

	//s.logger.Logger.Info("start apiServer")

	return s.server.ListenAndServe()
}

func NewServer(
	cfg *ServerConfig,
	db *store.Store,
	log *ServerLogger,
	) (newServer *Server) {

	newServer = &Server{
		config: cfg,
		router: NewRouter(),
		store: db,
		logger: log,
	}

	return
}

