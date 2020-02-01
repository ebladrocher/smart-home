package main

import (
	"smarthome/api/server"
	"smarthome/system/config"
)

func main() {
	start()
}

func start() {
	r, _ := config.ReadConfig()
	cfg := config.AppConfig{}
	cfg.ServerHost = r.ServerHost
	cfg.ServerPort = r.ServerPort
	cfg.Mode = r.Mode
	srv := server.NewServerConfig(&cfg)
	srvstrt := server.NewServer(srv)
	srvstrt.Start()
}




