package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Item models each individual artwork
type Item struct {
	ID          int64
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	Name        string
	Description string
	Source      string
	URLImage    string `json:"url_image" db:"url_image"`
}

// SelectItems gets all items from the DB
func SelectItems(db *sqlx.DB) ([]Item, error) {
	items := []Item{}
	err := db.Select(&items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

// SelectItemByID gets a single item by its `id` column.
func SelectItemByID(db *sqlx.DB, id int64) (Item, error) {
	item := Item{}
	err := db.Get(&item, "SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

// InsertItem inserts a given item into the DB
func InsertItem(db *sqlx.DB, item Item) (Item, error) {
	rows, err := db.NamedQuery(
		"INSERT INTO items (name, description, source, url_image)"+
			" VALUES (:name, :description, :source, :url_image)"+
			" RETURNING *",
		&item,
	)
	if err != nil {
		return Item{}, err
	}
	insertedItem := Item{}
	rows.Next()
	if err := rows.StructScan(&insertedItem); err != nil {
		return Item{}, err
	}
	return insertedItem, nil
}

// TODO: Delete item method

// TODO: Update item method
