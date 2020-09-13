package bd

import (
	"github.com/ZachIgarz/tweetApp/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la base de datos*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usuarioData, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {

		return usuarioData, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usuarioData.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuarioData, false
	}
	return usuarioData, true
}
