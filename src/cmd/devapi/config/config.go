package config

import "github.com/spf13/viper"

type Config struct {
	Port                      string `mapstructure:"PORT"`
	DBUrl                     string `mapstructure:"DB_URL"`
	GrpcAuthenticationService string `mapstructure:"grpc_authentication_service"`
	PersonalWebsiteServices   string `mapstructure:"personal_website_services"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./src/cmd/devapi/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
