package server

import (
	"context"
	"fmt"
	adaptors "github.com/ebladrocher/smarthome/adaptors"
	"github.com/ebladrocher/smarthome/system/store"
	postgres "github.com/ebladrocher/smarthome/system/store/postgrestore"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server ...
type Server struct {
	Config    *ServerConfig
	//Handlers  *handlers.Handlers
	server    *http.Server
	router    *mux.Router
	logger    *Logger
	repository store.UseCase
	isStarted bool
}

// Server ...
func (s *Server) Start() error {

	//handlers.NewHandlers(s.repository)
	SetHandlers(s.router, s.repository)
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: s.router,
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
	log *Logger,
) (*Server, error) {

	store := postgres.InitDB()
	userRepo := postgres.NewUserRepository(store)

	newServer := &Server{
		Config:   cfg,
		router:   mux.NewRouter(),
		repository: adaptors.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
		logger:   log,
	}

	///newServer.setControllers()

	return newServer, nil
}
