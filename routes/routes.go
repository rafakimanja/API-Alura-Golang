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
	r.HandleFunc("/personalidades", controllers.AllPersonas)
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersona)
	http.Handle("/", r) //registra o roteador

	log.Fatal(http.ListenAndServe(":8000", nil))
}
