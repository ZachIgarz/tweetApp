package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ZachIgarz/tweetApp/bd"
	"github.com/ZachIgarz/tweetApp/models"
)

/*ModificarPerfil modifica el perfil del usuario //metodo*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ah logrado modificar el registro del usuario "+err.Error(), 400)
	}

	w.WriteHeader(http.StatusCreated)

}
