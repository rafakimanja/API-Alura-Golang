package main

import (
	"api-golang/models"
	"api-golang/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor... Localhost:8000")
	routes.HandleRequest()
}
