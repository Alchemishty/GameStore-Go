package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Alchemishty/GameStore-Go/pkg/models"
	"github.com/Alchemishty/GameStore-Go/pkg/utils"
	"github.com/gorilla/mux"
)

var NewGame models.Game

func GetGame(w http.ResponseWriter, r *http.Request) {
	newGames := models.GetAllGames()
	res, _ := json.Marshal(newGames)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	ID, err := strconv.ParseInt(gameID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	gameDetails, _ := models.GetGameById(ID)
	res, _ := json.Marshal(gameDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	Game := &models.Game{}
	utils.ParseBody(r, Game)
	g := Game.CreateGame()
	res, _ := json.Marshal(g)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	ID, err := strconv.ParseInt(gameID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	game := models.DeleteGame(ID)
	res, _ := json.Marshal(game)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	var updateGame = &models.Game{}
	utils.ParseBody(r, updateGame)
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	ID, err := strconv.ParseInt(gameID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	gameDetails, db := models.GetGameById(ID)
	if updateGame.Name != "" {
		gameDetails.Name = updateGame.Name
	}
	if updateGame.Developer != "" {
		gameDetails.Developer = updateGame.Developer
	}
	if updateGame.Studio != "" {
		gameDetails.Studio = updateGame.Studio
	}
	db.Save(&gameDetails)
	res, _ := json.Marshal(gameDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
