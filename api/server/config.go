package server

import (
	"smarthome/system/config"
	)

type ServerConfig struct {
	Host string
	Port int
	RunMode string
}

func NewServerConfig(/*cfg *config.AppConfig*/) *ServerConfig {
	cfg, _ := config.ReadConfig()
	return &ServerConfig{
		Host: cfg.ServerHost,
		Port: cfg.ServerPort,
		RunMode: cfg.Mode,
	}
}