package controllers

import (
	"api-golang/database"
	"api-golang/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersona(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {

	var newP models.Personalidade

	json.NewDecoder(r.Body).Decode(&newP)
	database.DB.Create(&newP)

	json.NewEncoder(w).Encode(newP)
}
