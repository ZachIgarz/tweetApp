package routers

import (
	"encoding/json"
	"net/http"
	"github.com/ZachIgarz/tweetApp/bd"
	"github.com/ZachIgarz/tweetApp/models"
)

/*Registro es la funcion para crear en la BD el registro de usuarios*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	//El Body de un HTTP  requist es ubn stream! , solo se puede ler una ves
	//una ves leido , se destruye
	err =: json.NewDecoder(r.Body).Decode(&t)

	if err != nil{
		http.Error(w,"Error en los datos recibidos" +err.Error() , 400)
		return
	}
	//el largo de un string
	if len(t.Email) == 0{
		http.Error(w,"El Email de usuario es requerido" , 400)
		return
	}

	if len(t.Password) <6{
		http.Error(w, "Debe especificar una contrasena de almenos 6 caracteres" , 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true{
		http.Error(w, "Ya existe un usuario registrado con ese email" , 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario" + err.Error() , 400)
		return
	}

	if status == false{
		http.Error(w, "No se ah logrado insertar el registro del usuario"  , 400)
		return
	}
	
	w.WriteHeader(http.StatusCreated)



	
}
