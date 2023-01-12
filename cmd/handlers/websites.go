package handlers

import (
	"errors"
	"net/http"

	"github.com/Lelo88/EjemploGOWEB/pkg/response"
	"github.com/Lelo88/EjemploGOWEB/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var SECRET_KEY = "ABC"

func Get (c *gin.Context){
	
	//cuando procedemos a realizar el get, enviamos un token de seguridad por el header para realizarlo
	token := c.GetHeader("token")
	
	if token == SECRET_KEY {
		c.JSON(http.StatusUnauthorized, response.Err(nil))
		return
	}
	
	//request

	//process
	websites := services.Get()
	c.JSON(http.StatusAccepted, response.OK("success",websites))
}

//creamos el handler para crear el elemento

type Request struct{ //no tiene ID porque es lo que voy a ingresar. Al ingresarlo, se le designara el siguiente ID
	URL 		string `json:"url" validate:"required"` //en vez de pedir una validacion en una funcion lo hago con la etiqueta required
	Host 		string `json:"host" validate:"required"`
	Category 	string `json:"category" validate:"required"`
	Protected 	bool   `json:"protected"`
} //si agrego json con los campos vacios sin determinar que estos NO deben estarlos, se sumaran a la lista cuando realice POST


func Create(c *gin.Context){
	var request Request
	//verifico que lo que voy a pasar este bien, sino me va tirar un error de status
	if err:=c.ShouldBind(&request);err!=nil{
		c.JSON(http.StatusBadRequest, response.Err(err))
		return  
	}

	//compruebo la validacion 
	validate := validator.New()
	if err:=validate.Struct(request);err!=nil{
		c.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	//cuando se llama a la funcion create, esta nos devuelve la nueva web ingresada y un nil, o un error si lo que ingresamos ya se encuentran en nuestra base de datos
	website, err:= services.Create(request.URL, request.Host, request.Category, request.Protected) //asigno a esta variable 
	
	//si existe en nuestra base de datos, nos va a tirar error
	if err!= nil{
		if errors.Is(err,  services.ErrAlreadyExists) {
			c.JSON(http.StatusConflict, response.Err(err))
		}
	}

	c.JSON(http.StatusCreated, response.OK("success", website)) //enviamos un mensaje de que el recurso se creo de manera correcta
}