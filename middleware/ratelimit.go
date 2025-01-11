package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick_web_golang/log"
	"quick_web_golang/provider"
)

// PreMinuteLimit 通过计数器进行限流
func PreMinuteLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		limited, _, err := provider.Limiter.Minute.RateLimit(c.ClientIP(), 1)
		if err != nil {
			log.Error(err)
		}
		if limited {
			c.JSON(http.StatusBadRequest, gin.H{"message": "频繁请求"})
			c.Abort()
			return
		}
		c.Next()
	}
}
