package facade

import (
	"github.com/gin-gonic/gin"

	"service/templ_service"
)

func Templ(c *gin.Context) *templ_service.TemplService {
	return templ_service.New(c)
}
