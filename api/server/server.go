package server

import (
	"fmt"
	"github.com/ebladrocher/smarthome/api/server/controllers"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

// Server ...
type Server struct {
	Config *ServerConfig
	Controllers *controllers.Controllers
	server *http.Server
	router *mux.Router
	store  store.Store
	logger *Logger
}

// Server ...
func (s *Server) Start() {

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: s.router,
	}

	//s.logger.Logger.Info(
	//	"Server start at",
	//	zap.Any(s.Config.Host, s.Config.Port),
	//)


	if err := s.server.ListenAndServe(); err != nil {
		s.logger.Logger.Fatal(
			"server start error",
			zap.Error(err),
		)
	}
}

// NewServer ...
func NewServer(
	cfg *ServerConfig,
	store store.Store,
	log *Logger,
) (*Server, error) {

	newServer := &Server{
		Config: cfg,
		Controllers: controllers.NewControllers(),
		router: mux.NewRouter(),
		store:  store,
		logger: log,
	}

	//if err := newServer.store.SetStore();err != nil {
	//	return nil,err
	//}

	newServer.setControllers()

	return newServer, nil
}
