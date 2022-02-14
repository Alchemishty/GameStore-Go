package models

import (
	"github.com/Alchemishty/GameStore-Go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Game struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	Developer string `json:"developer"`
	Studio    string `json:"studio"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Game{})
}

func (g *Game) CreateGame() *Game {
	db.NewRecord(g)
	db.Create(&g)
	return g
}

func GetAllGames() []Game {
	var Games []Game
	db.Find(&Games)
	return Games
}

func GetGameById(Id int64) (*Game, *gorm.DB) {
	var getGame Game
	db := db.Where("ID=?", Id).Find(&getGame)
	return &getGame, db
}

func DeleteGame(Id int64) Game {
	var game Game
	db.Where("ID=?", Id).Delete(game)
	return game
}
