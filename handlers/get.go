// Package handlers implements basic get and post requests to work with the server
//
// GetName
// MainPage
// ToMain
// SignUp
// Auth
// SignOut
package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetName Return base JSON with message
func GetName(c *gin.Context) {
	var message string
	name := c.Param("name")
	message = fmt.Sprintf("Hello %s", name)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// MainPage Return base JSON with message
func MainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "welcome!",
	})
}

// ToMain Return base JSON with message
func ToMain(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/")
}

// SignUp Return base JSON with message
func SignUp(c *gin.Context) {
	c.JSON(200, "Send POST request to /auth/signup with params: name, email, password")

}

// Auth Return base JSON with message
func Auth(c *gin.Context) {
	c.JSON(200, "Send POST request to /auth/login with params: name, password")

}

// SignOut Return base JSON with message
func SignOut(c *gin.Context) {
	c.JSON(200, "Send POST request to /auth/logout")

}
