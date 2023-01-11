package main

import (
	"github.com/Lelo88/EjemploGOWEB/services"
	"github.com/gin-gonic/gin"
)

func main() {
	//server
	sv:= gin.Default()
	
	//router
	sv.GET("/websites", func(c *gin.Context){
		//request

		//router
		websites := services.Get()

		c.JSON(200, websites)
	})

	//star
	sv.Run(":8080")
}