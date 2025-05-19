package config

import (
	"github.com/spf13/viper"

	"config/assets_config"
)

const (
	Token = "token"
)

func Read() *viper.Viper {
	cfg := viper.New()
	assets_config.Read(cfg)
	return cfg
}
