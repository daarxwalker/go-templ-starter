package middleware

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"common/pkg/env"
	"common/pkg/facade"
	"common/pkg/view"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if env.Development() {
					fmt.Print(err)
				}
				status := c.Writer.Status()
				if status < 400 {
					status = http.StatusInternalServerError
				}
				facade.Templ(c).MustRender(status, view.Error(status, err.(error)))
				c.Abort()
				return
			}
		}()
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		status := c.Writer.Status()
		facade.Templ(c).MustRender(status, view.Error(status, err))
	}
}
