package config


import "github.com/spf13/viper"

func Init() error {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("configtoken")

	return viper.ReadInConfig()
}
