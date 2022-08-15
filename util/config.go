package util

import (
	"github.com/spf13/viper"
)

// hold all config of the app
// values are read by viper from config file or env vars
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or env vars
func LoadConfig(path string) (config Config, err error) {

	// for windows
	// viper.SetConfigFile(path + "\\app.env")

	// for unix
	viper.AddConfigPath(path)
	viper.SetConfigFile("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
