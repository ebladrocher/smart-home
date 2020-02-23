package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	"log"
)


func main() {
	start()
}

// Start ...
func start() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	cfg, _ :=config.ReadConfig()

	srv, err := server.NewServer(
		server.NewServerConfig(cfg),
		server.InitLogger(cfg),
		)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}
