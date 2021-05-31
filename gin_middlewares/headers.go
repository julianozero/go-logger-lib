package gin_middlewares

import (
	"github.com/gin-gonic/gin"
)

func ExtractHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		clientID := c.Request.Header.Get("X-NW-Client")
		userID := c.Request.Header.Get("X-NW-User")

		c.Set("requestId", requestID)
		c.Set("clientId", clientID)
		c.Set("userId", userID)

		c.Next()
	}
}
