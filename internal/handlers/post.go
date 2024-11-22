package handlers

import (
	"Server/pkg"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserReg struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserAuth struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var user UserReg
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while parsing data")
		return
	}
	hashString := PasswordHash(user.Password)

	conn, _ := pkg.Pool.Acquire(c)
	_, err := conn.Exec(c, `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`, user.Name, user.Email, hashString)
	if err != nil {
		pkg.BaseErrorHandler(c, err, "Error while insert user")
		return
	}

	session := sessions.Default(c)
	session.Set("name", user.Name)
	session.Set("logged_in", true)

	if err = session.Save(); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while saving session")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true", "message": "User registered successfully"})
}

func PasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func LogIn(c *gin.Context) {
	var user UserAuth
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while login user")
		return
	}

	DBUser := pkg.GetUserByName(c, user.Name)
	if DBUser == nil {
		return
	}
	if PasswordHash(user.Password) != DBUser.Password {
		pkg.BaseErrorHandler(c, nil, "Incorrect credentials.")
		return
	}

	session := sessions.Default(c)
	session.Set("name", user.Name)
	session.Set("logged_in", true)

	if err := session.Save(); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while saving session")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true", "message": "User logged in successfully"})

}

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while saving session")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
