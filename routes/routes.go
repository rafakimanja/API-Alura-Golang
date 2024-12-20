package routes

import (
	"api-golang/controllers"
	"api-golang/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersona).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersona).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.EditPersona).Methods("Put")

	http.Handle("/", r) //registra o roteador

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
