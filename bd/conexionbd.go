package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MongoCN=conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://admin:josueob73034147@168.235.69.64")

/*Conectar bd  es la funcion que me permite conectar la base de datos */

func conectarBD() *mongo.Client{
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(),nil)
	if err != nil{
		return 0
	}
	return 1
}