package config

type AppConfig struct {
	ServerHost string `json:"server_host"`
	ServerPort int `json:"server_port"`
	Mode string `json:"mode"`
}

//type RunMode string

/*const (
	DebugMode   = string("debug")
	ReleaseMode = string("release")
)*/
