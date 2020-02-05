package main

import (
	"smarthome/api/server"
	"smarthome/system/config"
	"smarthome/system/store"
)


func main() {
	start()
}

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
	dbConf := store.NewDbConfig(&cfg)
	storeConf := store.InitStore(dbConf)
	srv := server.NewServer(srvConf, storeConf, logConfig)
	srv.Start()
}




