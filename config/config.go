package config

import "github.com/spf13/viper"

type Config struct {
	PostgresUri string `mapstructure:"POSTGRES_URI"`
}

func NewConfig() (cfg *Config, err error) {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
