package config

type AppConfig struct {
	ServerHost string `json:"server_host"`
	ServerPort int    `json:"server_port"`
	Mode       string `json:"mode"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbName     string `json:"db_name"`
}
