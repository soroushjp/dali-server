package models

import "time"

// Item models each individual artwork
type Item struct {
	ID          int       `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	Source      string    `json:"source"`
	URLImage    string    `json:"url_image"`
}
