package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/br-jeff/go-gin-simple/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/players", controller.GetAllPlayers)
	r.GET("/:name", controller.GetByName)
	r.POST("/player", controller.CreatePlayer)
	r.GET("/player/:id", controller.GetPlayerByID)
	r.DELETE("/player/:id", controller.DeletePlayer)
	r.PATCH("/player/:id", controller.EditPlayer)
	r.Run(":8080")
}
