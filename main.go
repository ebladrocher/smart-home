package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	postgres "github.com/ebladrocher/smarthome/system/store/postgrestore"
)


func main() {
	start()
}

// Start ...
func start() {
	cfg, _ :=config.ReadConfig()
	dbUrl := postgres.NewDbConfig(cfg)

	srv, err := server.NewServer(
		server.NewServerConfig(cfg),
		dbUrl.ConnectionSting(),
		server.InitLogger(cfg),
		)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}
