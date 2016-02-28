package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soroushjp/dali-server/handlers"
)

func main() {
	r := gin.Default()

	itemsHandler := handlers.NewItemsHandler()

	r.GET("/items", itemsHandler.Index)
	r.POST("/items/:id", itemsHandler.Create)
	r.GET("/items/:id", itemsHandler.Read)
	r.PUT("/items", itemsHandler.Update)
	r.DELETE("/items", itemsHandler.Delete)

	r.Run(":3000")
}
