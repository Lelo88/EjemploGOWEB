package main

import (
	"net/http"

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
		websites := services.Get()

		c.JSON(http.StatusOK, response.OK("succed to websites",websites))
	})

	//start
	sv.Run(":8080")
}