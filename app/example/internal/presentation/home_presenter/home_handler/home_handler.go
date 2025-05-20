package home_handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"common/pkg/facade"
	"example/internal/infrastructure/repository/example_repository"
	"example/internal/presentation/home_presenter/home_view"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := example_repository.GetExample(c, facade.DB(c))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		facade.Templ(c).MustRender(http.StatusOK, home_view.Home(result == 1))
	}
}
