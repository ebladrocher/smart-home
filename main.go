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
	srvConf := server.NewServerConfig(&cfg)
	dbConf := store.NewDbConfig(&cfg)
	storeConf := store.Init(dbConf)
	srv := server.NewServer(srvConf,storeConf)
	srv.Start()
}




