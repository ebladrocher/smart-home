package main

import (
	"github.com/ebladrocher/smarthome/api/server"
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store/sqlstore"
)


func main() {
	start()
}

// Start ...
func start() {

	cfg, _ :=config.ReadConfig()

	thisDB, err  := sqlstore.ConnectToDB(
		sqlstore.NewDbConfig(cfg),
		)
	if err != nil {
		panic(err.Error())
	}

	srv, err := server.NewServer(
		server.NewServerConfig(cfg),
		sqlstore.NewStore(thisDB),
		server.InitLogger(cfg),
		)
	if err != nil {
		panic(err.Error())
	}
	srv.Start()
}




