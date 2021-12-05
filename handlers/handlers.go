package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/UrieelMM/clon-twitter-react-go/middlew"
	"github.com/UrieelMM/clon-twitter-react-go/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores setea el puerto, el handler y escucha el servidor
func Manejadores() {
	router := mux.NewRouter()

	//Rutas, se le debe pasar un middlew, una ruta y un metodo
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	//Esteblecer el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
