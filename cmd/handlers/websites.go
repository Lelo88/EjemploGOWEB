package handlers

import (
	"net/http"

	"github.com/Lelo88/EjemploGOWEB/pkg/response"
	"github.com/Lelo88/EjemploGOWEB/services"
	"github.com/gin-gonic/gin"
)

func Get (c *gin.Context){
	//request

	//router
	websites := services.Get()
	c.JSON(http.StatusAccepted, response.OK("success",websites))
}