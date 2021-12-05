package bd

import (
	"context"
	"github.com/UrieelMM/clon-twitter-react-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//ChequeoYaExisteUsuario recibe un email y revisa si ya existe en la base de datos
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
