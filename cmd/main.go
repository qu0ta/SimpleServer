package main

import (
	"Server/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {
	r := gin.New()
	file, err := os.Create("../logs/gin.log")
	if err != nil {
		fmt.Println("Error while creating logs file: ", err)
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

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
		auth.GET("/auth", handlers.SignUp)
		auth.GET("/auth", handlers.LogIn)
		auth.GET("/auth", handlers.LogOut)
	}
	r.GET("/", handlers.MainPage)
	r.Run()
}
