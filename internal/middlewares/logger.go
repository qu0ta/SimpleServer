package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		c.String(http.StatusOK, fmt.Sprintf("Latency: %v\n", latency))

	}
}
