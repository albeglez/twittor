package bd

import (
	"github.com/albeglez/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin realiza el chequeo de la login a la BD */
func TryLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := CheckExistingUser(email)
	if encontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
