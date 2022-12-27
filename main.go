package main

import (
	"log"

	"github.com/JosueOblitas/twittor/bd"
	"github.com/JosueOblitas/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0{
		log.Fatal("Sin Conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}