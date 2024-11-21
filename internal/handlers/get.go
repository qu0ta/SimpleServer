package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetName(c *gin.Context) {
	var message string
	name := c.Param("name")
	message = fmt.Sprintf("Hello %s", name)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
func MainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "welcome!",
	})
}

func ToMain(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/")
}

func SignUp(c *gin.Context) {
	c.JSON(200, "Send POST request to /auth/signup with params: name, email, password")

}

func LogIn(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/")

}

func LogOut(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/")

}
