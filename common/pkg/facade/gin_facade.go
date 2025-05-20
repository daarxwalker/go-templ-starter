package facade

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Gin(c context.Context) *gin.Context {
	ctx, ok := c.(*gin.Context)
	if !ok {
		panic("cannot type cast context.Context to *gin.Context")
	}
	return ctx
}
