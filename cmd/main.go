package main

import (
	"github.com/Lelo88/EjemploGOWEB/cmd/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	//server
	sv:= gin.Default()
	websites := sv.Group("/websites")
	//router
	websites.GET("/", handlers.Get)

	websites.POST("/", handlers.Create)
	
	//start
	sv.Run()
}