package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Item models each individual artwork
type Item struct {
	ID          uint64    `json:"id"`
	Created     time.Time `json:"created" db:"created"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Source      string    `json:"source" binding:"required"`
	URLImage    string    `json:"url_image" db:"url_image" binding:"required"`
}

// SelectItems gets all items from the DB
func SelectItems(db *sqlx.DB) ([]Item, error) {
	items := []Item{}
	err := db.Select(&items, "SELECT * FROM items")
	if err != nil {

		fmt.Println(err)

		return nil, err
	}
	return items, nil
}

// SelectItemByID gets a single item by its id.
func SelectItemByID(db *sqlx.DB, id uint64) (Item, error) {
	item := Item{}
	err := db.Get(&item, "SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

// InsertItem inserts the given item into the DB
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

// DeleteItem deletes an item given its id. Returns bool whether item was
// deleted, error otherwise.
func DeleteItem(db *sqlx.DB, id uint64) (bool, error) {
	result, err := db.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	numAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	// Since we are deleting by unique key id, we can guaranteed to get 1 or 0
	// for numAffected here.
	if numAffected == 0 {
		return false, nil
	}
	return true, nil
}

// UpdateItem updates the given item in the DB. Item is selected by id and all
// fields are updated. Returns bool whether item was updated, error otherwise.
func UpdateItem(db *sqlx.DB, item Item) (Item, error) {
	result, err := db.NamedExec(
		"UPDATE items SET"+
			" name = :name"+
			", description = :description"+
			", source = :source"+
			", url_image = :url_image WHERE id = :id",
		&item)
	if err != nil {
		return Item{}, err
	}
	numAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("hit here")
		return Item{}, err
	}
	// Since we are updating by unique key id, we can guaranteed to get 1 or 0
	// for numAffected here.
	if numAffected == 0 {
		return Item{}, sql.ErrNoRows
	}
	updatedItem, err := SelectItemByID(db, item.ID)
	if err != nil {
		return Item{}, err
	}
	return updatedItem, nil
}
