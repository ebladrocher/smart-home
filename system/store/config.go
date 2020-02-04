package store

import (
	"fmt"
	"smarthome/system/config"
)

type DbConfig struct {
	Name string
	Host string
	Port string
}

func (db *DbConfig) ConnectionSting() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", db.Host, db.Port, db.Name)
}

func NewDbConfig(cfg *config.AppConfig) *DbConfig  {
	return &DbConfig{
		Name: cfg.DbName,
		Host: cfg.DbHost,
		Port: cfg.DbPort,
	}
}

