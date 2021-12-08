package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func InitViper() *Config {
	v := viper.New()
	v.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(".")   // optionally look for config in the working directory
	// Find and read the config file
	if err := v.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	v.WatchConfig()
	return &Config{v}
}
