package routes

import (
	"api-golang/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersona).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersona).Methods("Delete")

	http.Handle("/", r) //registra o roteador

	log.Fatal(http.ListenAndServe(":8000", nil))
}
