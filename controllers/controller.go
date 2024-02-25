package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/br-jeff/go-gin-simple/models"
)

func GetAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Players)
}

func GetByName(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"Hello": "How are you ?" + name,
	})
}

func CreatePlayer(c *gin.Context) {

}
