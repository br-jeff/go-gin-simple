package main

import (
	"github.com/br-jeff/go-gin-simple/models"
	"github.com/br-jeff/go-gin-simple/routes"
)

func main() {
	models.Players = []models.Player{
		{Name: "Neymar Junior", Document: "213423234"},
		{Name: "Cristiano Ronaldo", Document: "6575656"},
	}
	routes.HandleRequest()
}
