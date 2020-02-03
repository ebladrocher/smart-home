package main

import (
	"smarthome/api/server"
)

func main() {
	start()
}

func start() {
	/*r, _ := config.ReadConfig()
	cfg := config.AppConfig{}
	cfg.ServerHost = r.ServerHost
	cfg.ServerPort = r.ServerPort
	cfg.Mode = r.Mode
	srv := server.NewServerConfig(&cfg)*/
	srv := server.NewServer()
	srv.Start()
}




