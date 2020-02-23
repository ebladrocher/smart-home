package server

import (
	"context"
	"fmt"
	"github.com/ebladrocher/smarthome/api/server/controllers"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server ...
type Server struct {
	Config      *ServerConfig
	Controllers *controllers.Controllers
	server      *http.Server
	router      *mux.Router
	store       store.Store
	logger      *Logger
	isStarted   bool
}

// Server ...
func (s *Server) Start() error {
	router := s.router

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: router,
	}

	//s.logger.Logger.Info(
	//	"Server start at",
	//	zap.Any(s.Config.Host, s.Config.Port),
	//)

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Logger.Fatal(
				"server start error",
				zap.Error(err),
			)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.server.Shutdown(ctx)
}

// Shutdown ...

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer ...
func NewServer(
	cfg *ServerConfig,
	store store.Store,
	log *Logger,
) (*Server, error) {

	newServer := &Server{
		Config:      cfg,
		Controllers: controllers.NewControllers(),
		router:      mux.NewRouter(),
		store:       store,
		logger:      log,
	}

	newServer.setControllers()

	return newServer, nil
}
