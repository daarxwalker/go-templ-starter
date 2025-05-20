package cache_config

import (
	"github.com/spf13/viper"
	
	"common/pkg/env"
)

const (
	Uri             = "cache-uri"
	Password        = "cache-password"
	DB              = "cache-db"
	DeleteBatchSize = "cache-delete-batch-size"
)

func Read(v *viper.Viper) {
	if env.Development() {
		v.Set(Uri, "project-dragonfly:6379")
		v.Set(Password, "")
	}
	v.Set(DB, 0)
	v.Set(DeleteBatchSize, 100)
}
