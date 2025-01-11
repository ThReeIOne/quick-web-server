package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick_web_golang/lib"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(lib.HeaderXToken)
		if token == "" {
			c.JSON(http.StatusUnauthorized, lib.Unauthorized)
			c.Abort()
			return
		}

		claims, err := lib.NewJWT().ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, lib.Unauthorized)
			c.Abort()
			return
		}
		c.Set(lib.GinCtxKeyClaims, claims)
		c.Next()
	}
}
