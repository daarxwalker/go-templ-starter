package main

import (
	"log"
	"net/http"

	"config"
	"facade"
	"middleware"
	"view"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := config.Read()
	router.Static("/static", "./public/static")
	router.Use(
		middleware.ServiceProvider(
			map[string]any{
				config.Token: cfg,
			},
		),
	)
	router.Use(middleware.Assets(cfg))
	router.GET(
		"/", func(c *gin.Context) {
			facade.Templ(c).MustRender(http.StatusOK, view.Home())
		},
	)
	log.Fatalln(router.Run())
}
