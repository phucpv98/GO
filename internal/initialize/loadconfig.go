package initialize

import (
	"fmt"
	"go-ecommerce/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
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
	// var config Config
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("Unable to decode into struct: %v", err)
	}
}
