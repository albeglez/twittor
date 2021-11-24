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
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.CheckBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.CheckBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.CheckBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
