package main

import (
	"strconv"

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
	r.PUT("/items/:id", itemsHandler.Update)
	r.DELETE("/items/:id", itemsHandler.Delete)

	r.Run(":" + strconv.Itoa(int(app.Env.Port)))
}
