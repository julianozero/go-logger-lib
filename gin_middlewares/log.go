package gin_middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianozero/go-logger-lib/nlog"
)

const (
	defaultMessage string = "Request log"
)

func Log(logger *nlog.NLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsedTime := time.Since(start)

		requestID := c.GetString("requestId")
		clientID := c.GetString("clientId")
		userID := c.GetString("userId")
		url := c.Request.URL.String()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		message := c.GetString("message")

		if message == "" {
			message = defaultMessage
		}

		status := c.Writer.Status()
		switch {
		case status >= 400 && status < 500:
			event := logger.Warn().
				TraceID(requestID).
				Org(clientID, userID).
				Req(url, method).
				Res(statusCode).
				ElapsedTime(elapsedTime)
			if len(c.Errors.Errors()) > 0 {
				event.Err(c.Errors.Last().Err)
			}
			event.Send(message)
		case status >= 500:
			event := logger.Error().
				TraceID(requestID).
				Org(clientID, userID).
				Req(url, method).
				Res(statusCode).
				ElapsedTime(elapsedTime)
			if len(c.Errors.Errors()) > 0 {
				event.Err(c.Errors.Last().Err)
			}
			event.Send(message)
		default:
			logger.Info().
				TraceID(requestID).
				Org(clientID, userID).
				Req(url, method).
				Res(statusCode).
				ElapsedTime(elapsedTime).
				Send(message)
		}
	}
}
