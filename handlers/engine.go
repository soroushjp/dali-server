package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soroushjp/dali-server/context"
)

// NewEngine returns a new gin engine with routes and middleware set.
func NewEngine(app *context.AppContext) *gin.Engine {
	r := gin.Default()

	itemsHandler := NewItemsHandler(app)
	r.GET("/items", itemsHandler.Index)
	r.POST("/items", itemsHandler.Create)
	r.GET("/items/:id", itemsHandler.Read)
	r.PUT("/items/:id", itemsHandler.Update)
	r.DELETE("/items/:id", itemsHandler.Delete)

	return r
}
