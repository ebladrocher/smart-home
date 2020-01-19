package server

import (
	"smarthome/system/config"
	)

type ServerConfig struct {
	Host string
	Port int
	RunMode config.RunMode
}

func NewServerConfig(cfg *config.AppConfig) *ServerConfig {
	return &ServerConfig{
		Host: cfg.ServerHost,
		Port: cfg.ServerPort,
		RunMode: cfg.Mode,
	}
}