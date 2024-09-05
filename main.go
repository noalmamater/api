package main

import (
	"api/database"
	"api/models"
	"api/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historia: "historia1"},
		{Id: 2, Nome: "nome2", Historia: "historia2"},
	}

	database.ConectaDB()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
