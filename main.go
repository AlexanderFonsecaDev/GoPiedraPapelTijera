package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

const (
	StaticRootPath string = "static"
)

func main() {
	// Crear un enrutador
	router := http.NewServeMux()
	port := ":8080"

	fs := http.FileServer(http.Dir(StaticRootPath))
	router.Handle("/"+StaticRootPath+"/", http.StripPrefix("/"+StaticRootPath+"/", fs))

	// Configuracion de rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new-game", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	log.Printf("Listening on port http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
