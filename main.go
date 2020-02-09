package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store"
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
	srv, err := server.NewServer(srvConf, storeConf, logConfig)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}




