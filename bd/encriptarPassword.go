package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword e ncripta la contrase;a enviada por parametro*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
