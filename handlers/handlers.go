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
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GrabaTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/optenerAvatar", middlew.ChequeoBD(routers.OptenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/optenerBanner", middlew.ChequeoBD(routers.OptenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(routers.AltaRelacion)).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(routers.BajaRelacion)).Methods("DELETE")

	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(routers.ConsultaRelacion)).Methods("GET")

	router.HandleFunc("/listaUsuario", middlew.ChequeoBD(routers.ListaUsuarios)).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(routers.LeoTweetsSeguidores)).Methods("GET")

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
