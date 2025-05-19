package facade

import (
	"context"

	"github.com/spf13/viper"

	"config"
)

func Config(c context.Context) *viper.Viper {
	cfg, ok := c.Value(config.Token).(*viper.Viper)
	if !ok {
		panic(config.Token + " not found in context")
	}
	return cfg
}
