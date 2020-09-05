package middlew

import (
	"net/http"

	"github.com/ZachIgarz/tweetApp/bd"
)

/*ChequeoBD  es el middleware que me permite conocer el estado de la base de datos!*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	/*Los middlewares tienen que recibir algo y devolver lo mismo*/

	return func(w http.ResponseWriter, r *http.Request) {

		//Pasamanos
		if bd.ChequeoConnection() == 0 {
			//Error conexion
			http.Error(w, "Conexion Perdida con La base de datos", 500)
			return
		}
		//No fallo ? paso todos los valores recibidos al proximo eslavon de la cadena!
		next.ServeHTTP(w, r)

	}

}
