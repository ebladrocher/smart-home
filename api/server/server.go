package server

import (
	"context"
	"fmt"
	adaptors "github.com/ebladrocher/smarthome/adaptors"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store"
	postgres "github.com/ebladrocher/smarthome/system/store/postgrestore"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server ...
type Server struct {
	Config *ServerConfig
	server *http.Server
	router *mux.Router
	logger *Logger
	authUC store.UseCase
	//isStarted  bool
}

// Server ...
func (s *Server) Start() error {

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: s.router,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Logger.Fatal(
				"server start error",
				zap.Error(err),
			)
		}
	}()

	//s.logger.Logger.Info(
	//	"Server start at",
	//	zap.Any(s.Config.Host, s.Config.Port),
	//)

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
	dbUrl string,
	log *Logger,
) (*Server, error) {

	db := postgres.InitDB(dbUrl)
	userRepo := postgres.NewUserRepository(db)

	token, err := config.TokenInit()
	if err != nil {
		return nil, err
	}

	newServer := &Server{
		Config: cfg,
		router: mux.NewRouter(),
		authUC: adaptors.NewAuthUseCase(
			userRepo,
			token.HashSalt,
			[]byte(token.SigningKey),
			token.Token,
		),
		logger: log,
	}
	SetHandlers(
		newServer.router,
		newServer.authUC,
	)

	return newServer, nil
}
