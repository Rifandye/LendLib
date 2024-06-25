package models

import (
	"log"
	"time"

	"example.com/LendLib/db"
)

type Borrow struct {
	ID int64
	BorrowDate time.Time
	ReturnDate time.Time
	User_Id int64
	Book_id int64
}

func (b *Borrow) CreateBorrow(bookId, userId int64) error {
	query := `
		INSERT INTO "Borrows" ("borrowDate", "returnDate", "user_id", "book_id")
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	
	borrowDate := time.Now()
	returnDate := borrowDate.AddDate(0, 0, 14)


	var borrowId int64

	err := db.DB.QueryRow(query, borrowDate, returnDate, userId, bookId).Scan(&borrowId)
	if err != nil {
		log.Printf("Error creating borrow: %v", err)
		return err
	}

	b.ID = borrowId


	updateQuery := `
		UPDATE "Books"
		SET "stock" = "stock" - 1
		WHERE "id" = $1
	`

	_, err = db.DB.Exec(updateQuery, bookId)

	if err != nil {
		log.Printf("Error updating book stock: %v", err)
	}

	return nil
}