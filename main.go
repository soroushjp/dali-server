package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soroushjp/dali-server/context"
	"github.com/soroushjp/dali-server/handlers"
)

func main() {
	r := gin.Default()

	app, err := context.NewAppContext()
	if err != nil {
		panic(err)
	}

	itemsHandler := handlers.NewItemsHandler(app)

	r.GET("/items", itemsHandler.Index)
	r.POST("/items", itemsHandler.Create)
	r.GET("/items/:id", itemsHandler.Read)
	r.PUT("/items", itemsHandler.Update)
	r.DELETE("/items", itemsHandler.Delete)

	r.Run(":3000")
}
