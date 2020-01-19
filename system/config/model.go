package config

type AppConfig struct {
	ServerHost string `json:"server_host"`
	ServerPort int `json:"server_port"`
	Mode RunMode `json:"mode"`
}

type RunMode string

const (
	DebugMode   = RunMode("debug")
	ReleaseMode = RunMode("release")
)
