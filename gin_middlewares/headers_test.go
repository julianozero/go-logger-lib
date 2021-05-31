package gin_middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	invoked := false

	engine.Use(ExtractHeaders())
	engine.GET("/test/middleware", func(c *gin.Context) {
		clientID, exists := c.Get("clientId")
		assert.True(t, exists)
		assert.Equal(t, "2", clientID)

		userID, exists := c.Get("userId")
		assert.True(t, exists)
		assert.Equal(t, "123", userID)

		requestID, exists := c.Get("requestId")
		assert.True(t, exists)
		assert.Equal(t, "abc", requestID)

		invoked = true
	})

	server := httptest.NewServer(engine)
	defer server.Close()

	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test/middleware", nil)
	req.Header.Set("X-NW-Client", "2")
	req.Header.Set("X-NW-User", "123")
	req.Header.Set("x-request-id", "abc")

	_, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.True(t, invoked)
}
