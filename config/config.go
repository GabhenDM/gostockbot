package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s \n", err)
	}

}
