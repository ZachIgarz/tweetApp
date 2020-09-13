package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ZachIgarz/tweetApp/middlew"
	"github.com/ZachIgarz/tweetApp/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	//captura el http, y el response y el require!
	router := mux.NewRouter()
	//END POIN
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

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
