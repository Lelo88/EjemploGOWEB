// controlador del modelo
package services

import (
	"errors"
	"fmt"
	"github.com/Lelo88/EjemploGOWEB/services/models"
)

var ErrAlreadyExists = errors.New("item already exists")

var websites = []models.Website{
	{ID: 1, URL: "https://www.google.com", Host: "google", Category: "search", Protected: false},
	{ID: 2, URL: "https://www.yahoo.com", Host: "yahoo", Category: "search", Protected: true},
	{ID: 3, URL: "https://www.mercadolibre.com", Host: "meli", Category: "e-commerce", Protected: false},
	{ID: 4, URL: "https://www.amazon.com", Host: "amazon", Category: "finance", Protected: true},
}

var lastID = 4

//read: devolvemos lo que esta previamente cargado
func Get() []models.Website{
	return websites
}

func GetByID(){
	
}

func ExistURL(url string) bool{
	for _, website := range websites {
        if website.URL == url {
            return true
        }
    }
    return false
}

//vamos a crear un nuevo recurso(elemento)
//con esta funcion podemos pasar un elemento a traves de un post
//tuvimos que cambiar el retorno de la funcion create
func Create(url, host, category string, protected bool) (models.Website, error){
	
	if ExistURL(url) {
		return models.Website{}, fmt.Errorf("%w, %s",ErrAlreadyExists, "url not unique")
	}
	
	lastID++
	website := models.Website{
		ID: lastID ,
		URL: url,
		Host: host,
        Category: category,
	    Protected: protected,
    }
	
	websites = append(websites, website)
    
	return website, nil
}