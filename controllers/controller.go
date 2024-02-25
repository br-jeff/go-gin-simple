package controller

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

func GetByName(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"Hello": "How are you ?" + name,
	})
}
