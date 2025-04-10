package middlewares

import (
	"github.com/gin-gonic/gin"
)

type CorsMiddleware interface {
	CheckCors() gin.HandlerFunc
}

type CorsMiddlewareImpl struct {
}

func NewCorsMiddleware() CorsMiddleware {
	return &CorsMiddlewareImpl{}
}

func (umi *CorsMiddlewareImpl) CheckCors() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

	}

}
