package server

import (
	"github.com/ebladrocher/smarthome/system/config"
	)

// ServerConfig ...
type ServerConfig struct {
	Host string
	Port int
	RunMode config.RunMode
}

// NewServerConfig ...
func NewServerConfig(cfg *config.AppConfig) *ServerConfig {
	return &ServerConfig{
		Host: cfg.ServerHost,
		Port: cfg.ServerPort,
		RunMode: cfg.Mode,
	}
}