package app

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	Debug      bool   `mapstructure:"DEBUG"`

	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

var Config = &AppConfig{}

func InitConfig() (_ *AppConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return Config, err
}
