package main

import (
	"git.target.com/DhimanGhosh/GoPrograming/vget/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("/", handlers.AppHandler)
		v1.GET("/health", handlers.HealthHandler)
		v1.GET("/getVersion/:appName/:key", handlers.GetVersionHandler)

	}
	r.Run()
}
