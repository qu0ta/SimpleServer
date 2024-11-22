// Package pkg contains utility functions
//
// This package includes utility functions:
// - GenerateSecret
// - BaseErrorHandler

package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

// GenerateSecret generates a random secret string.
//
// This function creates a byte slice of length 32 and fills it with random bytes using the rand.Read function.
// If an error occurs during the random number generation, it prints the error message and returns an empty string.
// Otherwise, it encodes the byte slice into a base64 string and returns it.
//
// Returns:
// - A string representing the generated secret.
func GenerateSecret() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error generating secret:", err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

// BaseErrorHandler handles the error response for the given context and error.
//
// Parameters:
// - c: the gin context object representing the HTTP request and response.
// - err: the error that occurred.
// - message: the error message to be logged and displayed.
//
// Returns:
// - None.
func BaseErrorHandler(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": err.Error()})
	slog.Error(fmt.Sprintf("%s: ", message), err.Error())
	log.Fatalln(fmt.Sprintf("%s: ", message) + err.Error())
}

// PasswordHash generates a SHA256 hash of the given password and returns it as a hexadecimal string.
//
// Parameters:
// - password: the password to be hashed (string).
//
// Returns:
// - A string representing the SHA256 hash of the password in hexadecimal format.

func PasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
