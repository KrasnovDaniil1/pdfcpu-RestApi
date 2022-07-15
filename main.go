package main

import (
	"restapi/handler"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Default())
	router.POST("/pagenum", handler.PageCountFileRestApi)
	// router.GET("/cards", handlers.GetCards)
	// router.DELETE("/cards/:id", handlers.DeleteCard)
	// router.GET("/tags", handlers.GetTags)

	router.Run("localhost:8080")
}
