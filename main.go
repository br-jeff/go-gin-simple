package main

import (
	"github.com/br-jeff/go-gin-simple/database"
	"github.com/br-jeff/go-gin-simple/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequest()
}
