package main

import (
	//"net/http"

	"github.com/Lelo88/EjemploGOWEB/pkg/response"
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

		c.JSON(200, response.OK("succed to websites",services.Get()))
	})

	//start
	sv.Run()
}