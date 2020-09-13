package middlew

import (
	"net/http"

	"github.com/ZachIgarz/tweetApp/routers"
)

/*ValidoJwt  permite validar el jwt que viene en la peticion*/
func ValidoJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
