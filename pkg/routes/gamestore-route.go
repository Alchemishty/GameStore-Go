package routes

import (
	"github.com/Alchemishty/GameStore-Go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterGameStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/game/", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/game/", controllers.GetGame).Methods("GET")
	router.HandleFunc("/game/{gameID}", controllers.GetGameById).Methods("GET")
	router.HandleFunc("/game/{gameID}", controllers.UpdateGame).Methods("PUT")
	router.HandleFunc("/game/{gameID}", controllers.DeleteGame).Methods("DELETE")

}
