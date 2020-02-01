package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"smarthome/system/store"
)

type Server struct {
	config *ServerConfig
	server *http.Server
	router *Router
	logger *logrus.Logger
	store *store.Store

}

func (s *Server) Start() error{
	s.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler: s.router.router,
	}

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.router.configureRouter()

	if err := s.store.ConfigureStore(); err != nil {
		return err
	}
	return s.server.ListenAndServe()

}


func NewServer(
	cfg *ServerConfig,
	) (newServer *Server) {

	newServer = &Server{
		config: cfg,
		logger: logrus.New(),
		router: New(),
		store: store.Init(),
	}

	return
}

func (s *Server) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.RunMode)
	if err != nil{
		return err
	}

	s.logger.SetLevel(level)
	return nil
}