package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-gin-simple/database"
	"github.com/br-jeff/go-gin-simple/models"
)

func GetAllPlayers(c *gin.Context) {
	var players []models.Player
	database.DB.Find(&players)

	c.JSON(http.StatusOK, players)

}

func GetByName(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"Hello": "How are you ?" + name,
	})
}

func CreatePlayer(c *gin.Context) {
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&player)

	c.JSON(http.StatusOK, player)
}
