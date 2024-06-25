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
	CreatedAt time.Time
	UpdatedAt time.Time
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

func GetBooks() ([]Book, error) {
	query := `
		SELECT * FROM "Books"
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		log.Printf("Error fetching books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []Book
	
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Name, &book.Description, &book.Image, &book.Author, &book.PublishedDate, &book.Stock, &book.CreatedAt, &book.UpdatedAt)

		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}


		books = append(books, book)
	}

	return books, nil
}

func GetBook(id int64) (*Book, error) {
	query := `
		SELECT * FROM "Books" WHERE id = $1
	`

	row := db.DB.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Name, &book.Description, &book.Image, &book.Author, &book.PublishedDate, &book.Stock, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}

	return &book, nil
}

func (b *Book) UpdateBook() error {
	query := `
		UPDATE "Books"
		SET "name" = $1, "description" = $2, "image" = $3, "author" = $4, "publishedDate" = $5, "stock" = $6
		WHERE id = $7
	`

	_, err := db.DB.Exec(query, b.Name, b.Description, b.Image, b.Author, b.PublishedDate, b.Stock, b.ID)

	return err
}