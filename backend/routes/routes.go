package routes

import (
	"tiny_first_steps/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")

    api.GET("/users/:id", controllers.GetUser)
    // api.POST("/users", controllers.CreateUser)
}
