package routers

import (
	"net/http"

	"github.com/ZachIgarz/tweetApp/bd"
)

/*EliminarTweet elimina un tweer*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id ", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar el tweety "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
