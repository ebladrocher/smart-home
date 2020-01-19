package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
	v1 "smarthome/api/server/controllers/v1"
	"smarthome/system/config"
)
var (
	log = logging.MustGetLogger("server")
)

type Server struct {
	Config *ServerConfig
	server *http.Server
	engine *gin.Engine
	Controllers *v1.Controllers
	logger *ServerLogger
}

func (s *Server) Start() {
	s.server = &http.Server{
		Addr: fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: s.engine,
	}

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		fmt.Println("listen: %s\n", err)
	}
}

func NewServer(
	cfg *ServerConfig,
	contr *v1.Controllers,
	) (newServer *Server) {

	logger := &ServerLogger{log}

	gin.DisableConsoleColor()
	gin.DefaultWriter = logger
	gin.DefaultErrorWriter = logger
	if cfg.RunMode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else if cfg.RunMode == config.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))

	newServer = &Server{
		Config: cfg,
		engine: engine,
		Controllers: contr,
	}

	newServer.setControllers()

	return
}