package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
)


func main() {
	start()
}

// Start ...
func start() {

	conf, _ :=config.ReadConfig()
	cfg := config.AppConfig{
		ServerHost:conf.ServerHost,
		ServerPort:conf.ServerPort,
		Mode:conf.Mode,
		DbHost:conf.DbHost,
		DbPort:conf.DbPort,
		DbName:conf.DbName,
	}
	logConfig := server.InitLogger(&cfg)
	srvConf := server.NewServerConfig(&cfg)
	srv, err := server.NewServer(srvConf, logConfig)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}




