package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store/db"
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

	tmp := db.NewDbConfig(&cfg)
	thisDB, err  := db.ConnectToDB(tmp)
	store := db.NewStore(thisDB)

	srvConf := server.NewServerConfig(&cfg)

	srv, err := server.NewServer(srvConf, store, logConfig)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}




