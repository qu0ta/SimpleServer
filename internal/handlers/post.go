package handlers

import (
	"Server/pkg"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
)

type UserForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var user UserForm
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Error("Error while parsing data: " + err.Error())
		log.Println("Error while parsing data: " + err.Error())
		c.JSON(200, gin.H{"success": false, "error": err})
		return
	}
	hash := sha256.Sum256([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	conn, _ := pkg.Pool.Acquire(c)
	_, err := conn.Exec(c, `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`, user.Name, user.Email, hashString)
	if err != nil {
		slog.Error("Error while insert user: ", err.Error())
		log.Fatalln("Error while insert user: " + err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("name", user.Name)
	session.Set("logged_in", true)

	if err = session.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": err.Error()})
		slog.Error("Error while saving session: ", err.Error())
		log.Fatalln("Error while saving session: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true", "message": "User registered successfully"})
}
