// Package handlers implements basic get and post requests to work with the server
//
// Register
// Auth
// LogIn
// LogOut
package handlers

import (
	"Server/pkg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserReg is a struct for user registration
type UserReg struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserAuth is a struct for user authentication
type UserAuth struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register handles the registration of a new user.
//
// It expects a JSON payload in the request body with the following fields:
// - Name: the name of the user (required)
// - Email: the email of the user (required)
// - Password: the password of the user (required)
//
// If the payload is successfully parsed, it hashes the password and inserts the user into the "users" table.
// It also sets the user's name and a "logged_in" flag in the session.
// Finally, it returns a JSON response with a success message.
//
// Parameters:
// - c: the gin context object representing the HTTP request and response.
//
// Returns:
// - None.
func Register(c *gin.Context) {
	var user UserReg
	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while parsing data")
		return
	}
	hashString := pkg.PasswordHash(user.Password)

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

// LogIn handles the login of a user.
//
// It expects a JSON payload in the request body with the following fields:
// - Name: the name of the user (required)
// - Password: the password of the user (required)
//
// It checks if the user exists in the database and if the provided password matches the stored hashed password.
// If the login is successful, it sets the user's name and a "logged_in" flag in the session and returns a JSON response with a success message.
//
// Parameters:
// - c: the gin context object representing the HTTP request and response.
//
// Returns:
// - None.
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
	if pkg.PasswordHash(user.Password) != DBUser.Password {
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

// LogOut handles the logout of a user.
//
// It clears the session and saves it. If there is an error while saving the session,
// it calls the BaseErrorHandler function with the appropriate error message.
// Finally, it returns a JSON response with a success message.
//
// Parameters:
// - c: the gin context object representing the HTTP request and response.
//
// Returns:
// - None.
func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		pkg.BaseErrorHandler(c, err, "Error while saving session")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
