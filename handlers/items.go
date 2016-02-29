package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soroushjp/dali-server/context"
	"github.com/soroushjp/dali-server/models"
)

// ErrorResponse holds response JSON for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(c *gin.Context, message string) {
	c.JSON(500, ErrorResponse{
		Error: message,
	})
}

// ItemsHandler provides methods for handling /items routes
type ItemsHandler struct {
	app *context.AppContext
}

// NewItemsHandler creates a new ItemsHandler.
func NewItemsHandler(app *context.AppContext) *ItemsHandler {
	return &ItemsHandler{
		app: app,
	}
}

// Index returns all items.
func (h *ItemsHandler) Index(c *gin.Context) {
	items, err := models.SelectItems(h.app.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(c, "no items found.")
		} else {
			respondWithError(c, "unexpected error: could not get items.")
		}
		return
	}
	c.JSON(200, items)
}

// Create creates a new item.
func (h *ItemsHandler) Create(c *gin.Context) {
	var reqItem models.Item
	if err := c.BindJSON(&reqItem); err != nil {
		respondWithError(c, "bad request body.")
		return
	}
	insertedItem, err := models.InsertItem(h.app.DB, reqItem)
	if err != nil {
		respondWithError(c, "unexpected error: could not create item.")
		panic(err)
	}
	c.JSON(200, insertedItem)
}

// Read reads an item.
func (h *ItemsHandler) Read(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"error": "bad query parameter: 'id'.",
		})
		return
	}
	item, err := models.SelectItemByID(h.app.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(c, "no item found with given id.")
		} else {
			respondWithError(c, "unexpected error: could not get item.")
		}
		return
	}
	c.JSON(200, item)
}

// Update updates an item.
func (h *ItemsHandler) Update(c *gin.Context) {

}

// Delete deletes an item.
func (h *ItemsHandler) Delete(c *gin.Context) {

}
