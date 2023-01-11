// controlador del modelo
package services

import "github.com/Lelo88/EjemploGOWEB/services/models"

var websites = []models.Website{
	{URL: "https://www.google.com", Host: "google", Category: "search", Protected: false},
	{URL: "https://www.yahoo.com", Host: "yahoo", Category: "search", Protected: true},
	{URL: "https://www.mercadolibre.com", Host: "meli", Category: "e-commerce", Protected: false},
	{URL: "https://www.amazon.com", Host: "amazon", Category: "finance", Protected: true}, 
}

var lastID = 4

//read: devolvemos lo que esta previamente cargado
func Get()[]models.Website{
	return websites
}

func GetByID(){

}

func Create(){

}