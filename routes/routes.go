package routes

import (
	"api/controllers"
	"api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonaliade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
