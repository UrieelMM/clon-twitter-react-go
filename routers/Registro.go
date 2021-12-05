package routers

import (
	"encoding/json"
	"github.com/UrieelMM/clon-twitter-react-go/bd"
	"github.com/UrieelMM/clon-twitter-react-go/models"
	"net/http"
)

//Registro es es el método que registra a los usarios en la base de datos
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	//Si hay un error en el registro
	if err != nil {
		http.Error(w, "Error en los datos recibidos."+err.Error(), 400)
		return
	}
	//Si no hay error validamos los datos recibidos
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es obligatorio.", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener mas de 6 caracteres.", 400)
		return
	}

	//Si el usuario (email) ya existe
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email.", 400)
		return
	}
	//Si se han pasado todas las validaciones se regitra el usuario en la base de datos
	_, status, err := bd.InsertoRegistro(t)
	//Si ocurre un error en la base de datos al registrar el usuario
	if err != nil {
		http.Error(w, "Ocurrió un error al registrar el usuario."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró registrar le usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
