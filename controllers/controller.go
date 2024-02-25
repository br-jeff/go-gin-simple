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

func GetPlayerByID(c *gin.Context) {
	var player models.Player
	id := c.Params.ByName("id")
	database.DB.First(&player, id)

	if player.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "We don't found the player",
		})
	}

	c.JSON(http.StatusOK, player)
}

func DeletePlayer(c *gin.Context) {
	var player models.Player

	id := c.Params.ByName("id")

	database.DB.Delete(&player, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "player was deleted",
	})

}

func EditPlayer(c *gin.Context) {
	var player models.Player
	id := c.Params.ByName("id")

	database.DB.First(&player, id)

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Model(&player).UpdateColumns(player)
	c.JSON(http.StatusOK, player)
}
