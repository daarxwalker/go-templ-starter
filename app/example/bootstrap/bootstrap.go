package bootstrap

import (
	"os"
	
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	
	"common/pkg/config"
	"common/pkg/env"
	"common/pkg/middleware"
	"common/pkg/service/cache_service"
	"common/pkg/service/database_service"
	"example/internal/presentation/home_presenter"
)

func Boot() error {
	if env.Production() {
		gin.SetMode(gin.ReleaseMode)
	}
	cfg := config.Read()
	app := gin.New()
	if env.Development() {
		app.Use(
			gin.LoggerWithConfig(
				gin.LoggerConfig{
					SkipPaths: []string{},
				},
			),
		)
	}
	app.Static("/static", "./public/static")
	app.Use(
		gzip.Gzip(
			gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/.well-known/appspecific/com.chrome.devtools.json"}),
		),
	)
	app.Use(
		middleware.ServiceLocator(
			map[string]any{
				config.Token:           cfg,
				cache_service.Token:    cache_service.New(cfg),
				database_service.Token: database_service.New(cfg),
			},
		),
	)
	app.Use(middleware.Assets(cfg))
	app.Use(middleware.ErrorHandler())
	home_presenter.Register(app)
	return app.Run("0.0.0.0:" + os.Getenv("APP_PORT"))
}
