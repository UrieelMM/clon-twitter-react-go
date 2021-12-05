package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword Funcion para ecnriptar la password
func EncriptarPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err

}
