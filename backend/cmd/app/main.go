package main

import (
	"log"
	"tiny_first_steps/config"
	"tiny_first_steps/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.Load()
    config.ConnectDatabase()

    r := gin.Default()

    routes.SetupRoutes(r)

    if err:= r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
