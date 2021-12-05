package bd

import (
	"context"
	"github.com/UrieelMM/clon-twitter-react-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Insertar registro del usuario en la base de datos
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	//Encriptar la password enviada por el usuario
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
