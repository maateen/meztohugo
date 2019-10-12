package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Database is an exportable type
type Database struct {
	Type     string `mapstructure: "type"`
	Hostname string `mapstructure: "hostname"`
	Port     string `mapstructure: "port"`
	DBName   string `mapstructure: "dbname"`
	Username string `mapstructure: "username"`
	Password string `mapstructure: "password"`
}

// Hugo ...
type Hugo struct {
	ContentDir        string
	ImageDir          string
	DefaultCoverImage string
}

// Config is an exportable type
type Config struct {
	Database Database `mapstructure: "database"`
	Hugo     Hugo     `mapstructure: "hugo"`
}

var config *Config

// GetConfig exports config settings
func GetConfig() *Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return config
}
