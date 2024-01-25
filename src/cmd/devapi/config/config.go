package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProfileSvcUrl string `mapstructure:"PROFILE_SVC_URL"`
	DBUser        string
	DBPassword    string
	DBUrl         string
	DBPort        string
	DBName        string
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./src/cmd/devapi/config/envs")

	scope := os.Getenv("SCOPE")
	log.Print("OS scope configuration:", scope)
	if len(scope) == 0 {
		scope = "dev"
	}
	log.Print("Selected configuration:", scope)
	viper.SetConfigName(scope)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	c.DBUser = readEnvironmentVariable("DB_USER")
	c.DBPassword = readEnvironmentVariable("DB_PASSWORD")
	c.DBUrl = readEnvironmentVariable("DB_URL")
	c.DBPort = readEnvironmentVariable("DB_PORT")
	c.DBName = readEnvironmentVariable("DB_NAME")

	return
}

func readEnvironmentVariable(envName string) string {
	envValue := os.Getenv(envName)
	if len(envValue) == 0 {
		panic(fmt.Sprintf("Environment Variable %s is not set", envName))
	}
	return envValue
}
