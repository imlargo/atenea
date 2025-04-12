package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imlargo/atenea/internal/ratelimiter"
	"github.com/imlargo/atenea/internal/responses"
)

func RateLimiterMiddleware(rl ratelimiter.Limiter, cfg ratelimiter.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cfg.Enabled {

			ip := c.ClientIP()

			if allow, retryAfter := rl.Allow(ip); !allow {
				message := "Rate limit exceeded. Try again in " + fmt.Sprintf("%.2f", retryAfter)
				responses.ErrorToManyRequests(c, message)
				return
			}
		}

		c.Next()
	}
}
