package main

import (
	"log"
	"net/http"

	"github.com/Alchemishty/GameStore-Go/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterGameStoreRoutes(r)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
