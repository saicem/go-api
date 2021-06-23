package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var config = new(Config)

type Config struct {
	MySQL   MySQL   `toml:"mysql"`
	Redis   Redis   `toml:"redis"`
	Session Session `toml:"session"`
}

func init() {
	viper.SetConfigName("dev" + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置已经被改变: %s", e.Name)
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() *Config {
	return config
}
