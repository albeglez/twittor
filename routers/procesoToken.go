package routers

import (
	"errors"

	"github.com/albeglez/twittor/bd"
	"github.com/albeglez/twittor/models"
	"github.com/dgrijalva/jwt-go"
)

/*Email valor de Email usado en todos los Endpoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los Endpoints	 */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Mimamamemima")
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.CheckExistingUser(claims.Email)
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
