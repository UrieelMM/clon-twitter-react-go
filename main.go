package main

import (
	"log"

	"github.com/UrieelMM/clon-twitter-react-go/bd"
	"github.com/UrieelMM/clon-twitter-react-go/handlers"
)

func main() {

	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
