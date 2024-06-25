package models

import (
	"log"
	"time"

	"example.com/LendLib/db"
)


type Book struct {
	ID int64
	Name string
	Description string
	Image string
	Author string
	PublishedDate time.Time
	Stock int
}

func (b *Book) CreateBook() error {
	query := `
		INSERT INTO "Books" ("name", "description", "image", "author", "publishedDate", "stock")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var bookId int64

	err := db.DB.QueryRow(query, b.Name, b.Description, b.Image, b.Author, b.PublishedDate, b.Stock ).Scan(&bookId)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	b.ID = bookId
	return nil
}
