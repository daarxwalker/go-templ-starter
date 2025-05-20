package home_presenter

import (
	"github.com/gin-gonic/gin"
	
	"example/internal/presentation/home_presenter/home_handler"
)

func Register(app gin.IRouter) {
	app.GET("/", home_handler.Home())
}
