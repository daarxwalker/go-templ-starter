package middleware

import (
	"github.com/gin-gonic/gin"
)

func ServiceLocator(services map[string]any) gin.HandlerFunc {
	return func(c *gin.Context) {
		for serviceToken, service := range services {
			c.Set(serviceToken, service)
		}
		c.Next()
	}
}
