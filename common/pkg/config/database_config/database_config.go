package database_config

import (
	"github.com/spf13/viper"
	
	"common/pkg/env"
)

const (
	Uri = "database-uri"
)

func Read(v *viper.Viper) {
	if env.Development() {
		v.Set(Uri, "postgresql://project:project@project-postgres:5432/project")
	}
}
