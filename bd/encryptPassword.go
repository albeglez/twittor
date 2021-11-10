package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncryptPassword es la rutina que me permite encriptar el password */
func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
