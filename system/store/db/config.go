package db

import (
	"fmt"
	"github.com/ebladrocher/smarthome/system/config"
)

// ConfigDB ...
type ConfigDB struct {
	Name string
	Host string
	Port string
}

// ConnectionString ...
func (db *ConfigDB) ConnectionSting() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", db.Host, db.Port, db.Name)
}

// NewDbConfig ,,,
func NewDbConfig(cfg *config.AppConfig) *ConfigDB  {
	return &ConfigDB{
		Name: cfg.DbName,
		Host: cfg.DbHost,
		Port: cfg.DbPort,
	}
}

