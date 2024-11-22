package main

import (
	"Server/internal/handlers"
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
		greetings.GET("/:name", handlers.GetName)
		greetings.GET("/", handlers.ToMain)
	}

	auth := r.Group("/auth")
	{
		auth.GET("/signup", handlers.SignUp)
		auth.POST("/signup", handlers.Register)
		auth.GET("/login", handlers.Auth)
		auth.POST("/login", handlers.LogIn)
		auth.GET("/logout", handlers.SignOut)
		auth.POST("/logout", handlers.LogOut)
	}
	r.GET("/", handlers.MainPage)
	r.GET("/session-data", func(c *gin.Context) {
		session := sessions.Default(c)

		c.JSON(200, gin.H{
			"name":      session.Get("name"),
			"logged_in": session.Get("logged_in"),
		})
	})
	r.Run()
}
