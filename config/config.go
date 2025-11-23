package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUri string `mapstructure:"DATABASE_URL"`
}

func NewConfig() (cfg *Config, err error) {

	viper.AddConfigPath(".")
	if os.Getenv("ENVIRONMENT") == "testing" {
		viper.SetConfigName(".env.testing")
	} else {
		viper.SetConfigName(".env")
	}

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
