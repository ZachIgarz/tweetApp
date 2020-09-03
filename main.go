package main

import (
	"log"

	"github.com/ZachIgarz/tweetApp/bd"
	"github.com/ZachIgarz/tweetApp/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin Conexion a la BD")
		return
	}

	handlers.Manejadores()

}
