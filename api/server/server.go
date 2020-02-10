package server

import (
	"database/sql"
	"fmt"
	"github.com/ebladrocher/smarthome/api/server/controllers"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store"
	"github.com/ebladrocher/smarthome/system/store/db"
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
	//store store.Store,
	log *Logger,
) (*Server, error) {
	tmp := config.AppConfig{
		DbHost:"localhost",
		DbPort:"5432",
		DbName: "smarthome_test",
	}
	thisConfig := db.NewDbConfig(&tmp)

	thisDB, err := newDB(thisConfig.ConnectionSting())
	if err != nil {
		return nil, err
	}

	defer thisDB.Close()
	store := db.NewStore(thisDB)

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

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
