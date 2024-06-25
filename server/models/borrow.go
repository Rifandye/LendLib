package models

import (
	"errors"
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
	//Validating user vould not borrowing the same book more than once
	var count int

	checkQuery := `
		SELECT COUNT(*) FROM "Borrows"
		WHERE "user_id" = $1 AND "book_id" = $2
	`

	err := db.DB.QueryRow(checkQuery, userId, bookId).Scan(&count)

    if err != nil {
        log.Printf("Error checking existing borrow: %v", err)
        return err
    }

    if count > 0 {
        return errors.New("user already borrowed this book")
    }



	//Inserting borrow data into database
	insertQuery := `
		INSERT INTO "Borrows" ("borrowDate", "returnDate", "user_id", "book_id")
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	
	borrowDate := time.Now()
	returnDate := borrowDate.AddDate(0, 0, 14)


	var borrowId int64

	err = db.DB.QueryRow(insertQuery, borrowDate, returnDate, userId, bookId).Scan(&borrowId)
	if err != nil {
		log.Printf("Error creating borrow: %v", err)
		return err
	}

	b.ID = borrowId


	//Updating stock of the book after successfully inserting borrow data
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