package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	connStr := os.Getenv("PG_CONNSTR")


	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	createTables()

	fmt.Println("Connected To Database")
}

func createTables() {
	createUserTables := `
	CREATE TABLE IF NOT EXISTS "Users" (
		"id" SERIAL PRIMARY KEY,
		"firstName" VARCHAR NOT NULL,
		"lastName" VARCHAR NOT NULL,
		"email" VARCHAR NOT NULL UNIQUE,
		"password" VARCHAR NOT NULL,
		"role" VARCHAR NOT NULL DEFAULT 'User',
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		"updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createUserTables)

	if err != nil {
		log.Fatalf("Error creating Users table: %v", err)
	}


	createBookTables := `
		CREATE TABLE IF NOT EXISTS "Books" (
		"id" SERIAL PRIMARY KEY,
		"name" VARCHAR NOT NULL,
		"description" TEXT NOT NULL,
		"image" VARCHAR NOT NULL,
		"author" VARCHAR NOT NULL,
		"publishedDate" TIMESTAMP NOT NULL,
		"stock" INTEGER,
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		"updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = DB.Exec(createBookTables)

	if err != nil {
		log.Fatalf("Error creating Books table: %v", err)
	}


	createBorrowTables := `
		CREATE TABLE IF NOT EXISTS "Borrows" (
		"id" SERIAL PRIMARY KEY,
		"borrowDate" TIMESTAMP NOT NULL,
		"returnDate" TIMESTAMP NOT NULL,
		"status" VARCHAR NOT NULL DEFAULT 'Lended',
		"user_id" INTEGER REFERENCES "Users"("id"),
		"book_id" INTEGER REFERENCES "Books"("id"),
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		"updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = DB.Exec(createBorrowTables)

	if err != nil {
		log.Fatalf("Error creating Borrows table: %v", err)
	}

}