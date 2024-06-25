package models

import (
	"errors"
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
	Role string 
}

type UserCredentials struct {
	ID        int64  
	FirstName string
	LastName  string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Role	  string
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

func (u *UserCredentials) ValidateCredentials() error {
	query := `
	SELECT "id", "firstName", "lastName", "email", "password", "role"
	FROM "Users" 
	WHERE "email" = $1
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedUser User

	err := row.Scan(&retrievedUser.ID, &retrievedUser.FirstName, &retrievedUser.LastName, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Role)

	if err != nil {
		return errors.New("invalid credentials")
	}

	u.ID = retrievedUser.ID
	u.FirstName = retrievedUser.FirstName
	u.LastName = retrievedUser.LastName
	u.Email = retrievedUser.Email
	u.Role = retrievedUser.Role

	validatePassword := utils.ComparePassword(u.Password, retrievedUser.Password)

	if !validatePassword {
		return errors.New("invalid credentials")
	}
	
	return nil
}