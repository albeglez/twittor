package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/albeglez/twittor/middlew"
	"github.com/albeglez/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
