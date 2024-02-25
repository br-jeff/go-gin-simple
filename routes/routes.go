package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/br-jeff/go-gin-simple/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/players", controller.GetAllPlayers)
	r.GET("/:name", controller.GetByName)
	r.Run(":8080")
}
