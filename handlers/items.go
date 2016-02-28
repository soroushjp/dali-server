package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soroushjp/dali-server/models"
)

// DEBUG ONLY REMOVE
var myItems = []models.Item{
	{
		ID:          1,
		CreatedDate: time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Name:        "The Persistence of Memory",
		Description: "By Salvador Dali (1931)",
		Slug:        "the-persistence-of-memory",
		Source:      "custom",
		URLImage:    "http://uploads5.wikiart.org/images/salvador-dali/the-persistence-of-memory-1931.jpg",
	},
	{
		ID:          2,
		CreatedDate: time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Name:        "Ballerina in a Death's Head",
		Description: "By Salvador Dali (1939)",
		Slug:        "ballerina-in-a-deaths-head",
		Source:      "custom",
		URLImage:    "http://uploads2.wikiart.org/images/salvador-dali/ballerina-in-a-death-s-head.jpg",
	},
}

// ItemsHandler provides methods for handling /items routes
type ItemsHandler struct{}

// NewItemsHandler creates a new ItemsHandler.
func NewItemsHandler() *ItemsHandler {
	return &ItemsHandler{}
}

// Index returns all items.
func (i *ItemsHandler) Index(c *gin.Context) {
	c.JSON(200, myItems)
}

// Create creates a new item.
func (i *ItemsHandler) Create(c *gin.Context) {

}

// Read reads an item.
func (i *ItemsHandler) Read(c *gin.Context) {

}

// Update updates an item.
func (i *ItemsHandler) Update(c *gin.Context) {

}

// Delete deletes an item.
func (i *ItemsHandler) Delete(c *gin.Context) {

}
