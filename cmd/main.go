package main

import (
	"github.com/Lelo88/EjemploGOWEB/cmd/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	//server
	sv:= gin.Default()
	
	//router
	sv.GET("/websites", handlers.Get)

	//start
	sv.Run()
}