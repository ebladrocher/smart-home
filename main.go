package main

import (
	"database/sql"
	"smarthome/api/server"
	v1 "smarthome/api/server/controllers/v1"
	"smarthome/system/config"
)

func main() {
	start()
	connStr := "user=postgres password=Pidor798718 dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)

}

func start() {
	r, _ := config.ReadConfig()
	cfg := config.AppConfig{}
	cfg.ServerHost = r.ServerHost
	cfg.ServerPort = r.ServerPort
	srv := server.NewServerConfig(&cfg)
	contr := v1.Controllers{
		Index:v1.NewControllerIndex(),
	}
	srvstrt := server.NewServer(srv, &contr)
	srvstrt.Start()
}




