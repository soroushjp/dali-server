package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/soroushjp/dali-server/context"
	"github.com/soroushjp/dali-server/models"
)

// ErrorResponse holds response JSON for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse holds response JSON for message responses.
type MessageResponse struct {
	Message string `json:"message"`
}

func respondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Error: message,
	})
}

func respondWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, MessageResponse{
		Message: message,
	})
}

func respondWithNoMatchingItem(c *gin.Context) {
	respondWithError(c, http.StatusNotFound, "no item found for given id.")
}

func respondWithBadQueryParam(c *gin.Context, param string) {
	respondWithError(c, http.StatusInternalServerError, "bad query parameter: '"+param+"'.")
}

func respondWithInvalidRequest(c *gin.Context) {
	respondWithError(c, http.StatusInternalServerError, "invalid request body.")
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
			respondWithError(c, http.StatusNotFound, "no items found.")
			return
		}
		respondWithError(c, http.StatusInternalServerError, "unexpected error: could not get items.")
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, items)
}

// Create creates a new item.
func (h *ItemsHandler) Create(c *gin.Context) {
	var reqItem models.Item
	if err := c.BindJSON(&reqItem); err != nil {
		respondWithInvalidRequest(c)
		return
	}
	insertedItem, err := models.InsertItem(h.app.DB, reqItem)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "unexpected error: could not create item.")
		c.Error(err)
	}
	c.JSON(http.StatusOK, insertedItem)
}

// Read reads an item.
func (h *ItemsHandler) Read(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		respondWithBadQueryParam(c, "id")
		return
	}
	item, err := models.SelectItemByID(h.app.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithNoMatchingItem(c)
			return
		}
		respondWithError(c, http.StatusInternalServerError, "unexpected error: could not get item.")
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, item)
}

// Update updates an item.
func (h *ItemsHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		respondWithBadQueryParam(c, "id")
		return
	}
	reqItem := models.Item{}
	if err = c.BindJSON(&reqItem); err != nil {
		respondWithInvalidRequest(c)
		return
	}
	reqItem.ID = id
	updatedItem, err := models.UpdateItem(h.app.DB, reqItem)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithNoMatchingItem(c)
			return
		}
		respondWithError(c, http.StatusInternalServerError, "unexpected error: could not update item.")
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, updatedItem)
}

// Delete deletes an item.
func (h *ItemsHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		respondWithBadQueryParam(c, "id")
		return
	}
	deleted, err := models.DeleteItem(h.app.DB, id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "unexpected error: could not delete item.")
		c.Error(err)
		return
	}
	if !deleted {
		respondWithNoMatchingItem(c)
		return
	}
	respondWithMessage(c, "item deleted.")
}
