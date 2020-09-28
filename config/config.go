package config

import (
	"fmt"

	"github.com/spf13/viper"
)


// AppConfig Represetn app config
type AppConfig struct {
	Name string 
	Global bool 
	Test   float64 
}

// Read Reads config from file
func Read() AppConfig {
	var result AppConfig
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error read config %v", err)
	}

	name := viper.GetString("name")

	fmt.Println(name)

	if err := viper.Unmarshal(&result); err != nil {
		fmt.Printf("Error load config to AppConfig %v\n", err)
	}

	return result
}