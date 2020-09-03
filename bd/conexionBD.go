package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC es el objeto de conexion a la bd*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://zach:********@cluster0.4r3c4.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la base de datos*/
func ConectarBD() *mongo.Client {
	//Contexto entorno de ejecucion TODO es default
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//Ping a la bases de datos
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Conexion Exitosa con la BD")
	return client

}

/*ChequeoConnection es el ping a la base de datos */
func ChequeoConnection() int {
	error := MongoCN.Ping(context.TODO(), nil)
	if error != nil {
		return 0
	}
	return 1
}
