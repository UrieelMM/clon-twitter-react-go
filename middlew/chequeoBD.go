package middlew

import (
	"net/http"

	"github.com/UrieelMM/clon-twitter-react-go/bd"
)

// Los middlew deben devolver lo mismo que reciben
//r es la petición del usuario y w la respuesta del servidor
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
