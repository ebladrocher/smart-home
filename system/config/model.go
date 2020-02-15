package config

// AppConfig ...
type AppConfig struct {
	ServerHost string `json:"server_host"`
	ServerPort int    `json:"server_port"`
	Mode       RunMode `json:"mode"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbName     string `json:"db_name"`
}

// RunMode ...
type RunMode string

const (
	// DebugMode ...
	DebugMode   = RunMode("debug")
	// RunMode ...
	ReleaseMode = RunMode("release")
)

