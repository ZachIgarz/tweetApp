package routers

import (
	"errors"
	"strings"

	"github.com/ZachIgarz/tweetApp/bd"
	"github.com/ZachIgarz/tweetApp/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de email usado en todos los endpoint*/
var Email string

/*IDUsuario es el id devuelto del modelo que se usara en todos los endpoints*/
var IDUsuario string

/*ProcesoToken  proceso el token para estraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo")
	claims := &models.Claim{}
	//Un split D:
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
