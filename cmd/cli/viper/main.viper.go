package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database []struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config
	viper.SetConfigName("local")      // ten file
	viper.SetConfigType("yaml")

	// read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// read server configuration
	fmt.Println("Database Host:", viper.GetString("database.host"))
	fmt.Println("Server Port:", viper.GetInt("server.port"))

	// configure structure
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Unable to decode into struct: %v", err)
	}

	fmt.Println("Config struct: ", config.Server.Port)
}
