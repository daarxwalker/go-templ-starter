package config

import (
	"github.com/spf13/viper"
	
	"common/pkg/config/assets_config"
	"common/pkg/config/cache_config"
	"common/pkg/config/database_config"
)

const (
	Token = "token"
)

func Read() *viper.Viper {
	cfg := viper.New()
	assets_config.Read(cfg)
	cache_config.Read(cfg)
	database_config.Read(cfg)
	return cfg
}
