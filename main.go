package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"name": "Neymar",
	})
}

func main() {
	r := gin.Default()

	r.GET("/players", GetAllPlayers)

	r.Run(":8080")
}
