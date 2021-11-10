package main

import (
	"log"

	"github.com/albeglez/twittor/bd"
	"github.com/albeglez/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
