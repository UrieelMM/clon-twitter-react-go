package bd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexión a la base de datos
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://urieelmm:trance76@twitterclon.eseyp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//ConectarDB es la función que me permite conectar a la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	fmt.Println("Conexión exitosa con la DB")
	return client
}

//CheckConnection es el Ping a la DB
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
