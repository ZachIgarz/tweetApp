package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	//captura el http, y el response y el require!
	router := mux.NewRouter()
	//verifica variable de entorno del sistema
	PORT := os.Getenv("PORT")
	if PORT == "" {

		PORT = "8080"
	}
	//Permisos a cualquiera
	handler := cors.AllowAll().Handler(router)
	//Escucha el puerto para ver las peticiones
	//Agrega el puerto a la url
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
