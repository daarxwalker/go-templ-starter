package database_service

import (
	"context"
	"log"
	"time"
	
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	
	"common/pkg/config/database_config"
)

func New(cfg *viper.Viper) Client {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	config, err := pgxpool.ParseConfig(cfg.GetString(database_config.Uri))
	if err != nil {
		log.Fatalf("unable to parse database config: %v\n", err)
	}
	pool, err := pgxpool.NewWithConfig(c, config)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	if pingErr := pool.Ping(c); pingErr != nil {
		log.Fatalf("ping to database failed: %v\n", err)
	} else {
		log.Println("connected and pinged to database")
	}
	return pool
}
