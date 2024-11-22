package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

func GenerateSecret() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error generating secret:", err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func BaseErrorHandler(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": err.Error()})
	slog.Error(fmt.Sprintf("%s: ", message), err.Error())
	log.Fatalln(fmt.Sprintf("%s: ", message) + err.Error())
}
