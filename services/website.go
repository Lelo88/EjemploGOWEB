// controlador del modelo
package services

import "github.com/Lelo88/EjemploGOWEB/services/models"

var websites = []models.Website{
	{ID: 1, URL: "https://www.google.com", Host: "google", Category: "search", Protected: false},
	{ID: 2, URL: "https://www.yahoo.com", Host: "yahoo", Category: "search", Protected: true},
	{ID: 3, URL: "https://www.mercadolibre.com", Host: "meli", Category: "e-commerce", Protected: false},
	{ID: 4, URL: "https://www.amazon.com", Host: "amazon", Category: "finance", Protected: true},
}

//var lastID = 4

//read: devolvemos lo que esta previamente cargado
func Get() []models.Website{
	return websites
}

func GetByID(){
	
}

func Create(){

}