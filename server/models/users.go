package models

import (
	"log"

	"example.com/LendLib/db"
	"example.com/LendLib/utils"
)


type User struct {
	ID        int64
	FirstName string 	`binding:"required"` 
	LastName string 	`binding:"required"` 
	Email string 		`binding:"required"` 
	Password string 	`binding:"required"` 
}


func (u *User) CreateUser() error {
	query := `
		INSERT INTO "Users" ("firstName", "lastName", "email", "password")
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var userId int64

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.FirstName, u.LastName, u.Email, hashedPassword).Scan(&userId)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	u.ID = userId
	return nil
}	