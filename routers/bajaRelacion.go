package routers

import (
	"net/http"

	"github.com/ZachIgarz/tweetApp/bd"
	"github.com/ZachIgarz/tweetApp/models"
)

/*BajaRelacion da de baja una relacion de usuarios borrado*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro ID es Obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se logro borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
