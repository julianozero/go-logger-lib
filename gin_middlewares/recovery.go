package gin_middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianozero/go-logger-lib/nlog"
)

func Recovery(logger *nlog.NLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error().
					Req(c.Request.URL.String(), c.Request.Method).
					Res(http.StatusInternalServerError).
					Err(err.(error)).
					Send("recovering from panic")

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
