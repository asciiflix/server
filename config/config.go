package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Db       string `mapstructure:"POSTGRES_DB"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
}

type APIConfig struct {
	Port   int    `mapstructure:"API_PORT"`
	JWTKey string `mapstructure:"JWT_PRIVATE_KEY"`
}

var Version = "development"
var Database DBConfig
var ApiConfig APIConfig

func GetConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err_db := viper.Unmarshal(&Database)
	err_api := viper.Unmarshal(&ApiConfig)

	if err_db != nil || err_api != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

}
