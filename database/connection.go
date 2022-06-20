package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	database, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic("Could not connect to database.db!")
	}
	log.Println("Database connected...!!!")

	// Create users table
	statement, _ := database.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			userID INTEGER PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			firstName VARCHAR(255),
			lastName VARCHAR(255),
			address VARCHAR(255),
			city VARCHAR(255),
			state VARCHAR(255),
			zipCode INT,
			phoneNumber VARCHAR(255),
			pin VARCHAR(9) NOT NULL
		);
	`)
	statement.Exec()

	// Assign DB for database connection so we can user in other packages
	DB = database
}
