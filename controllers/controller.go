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
