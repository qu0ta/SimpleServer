// Package main contains the entry point of the program and the main api handles
//
// This package includes main functions to interact with the server:
package main

import (
	handlers2 "Server/handlers"
	"Server/logging"
	"Server/pkg"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	pkg.InitDB()
	logging.InitLogging()
	r := gin.New()

	secret := pkg.GenerateSecret()
	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s  %d %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	greetings := r.Group("/greetings")
	{
		greetings.GET("/:name", handlers2.GetName)
		greetings.GET("/", handlers2.ToMain)
	}

	auth := r.Group("/auth")
	{
		auth.GET("/signup", handlers2.SignUp)
		auth.POST("/signup", handlers2.Register)
		auth.GET("/login", handlers2.Auth)
		auth.POST("/login", handlers2.LogIn)
		auth.GET("/logout", handlers2.SignOut)
		auth.POST("/logout", handlers2.LogOut)
	}
	r.GET("/", handlers2.MainPage)
	r.GET("/session-data", func(c *gin.Context) {
		session := sessions.Default(c)

		c.JSON(200, gin.H{
			"name":      session.Get("name"),
			"logged_in": session.Get("logged_in"),
		})
	})
	r.Run()
}
