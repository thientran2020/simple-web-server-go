package database

import (
	"database/sql"
	"log"
)

func Connect() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic("Could not connect to database.db!")
	}
	log.Println("Database connected...!!!")

	// Create users table
	statement, _ := db.Prepare(`
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

	// rows, _ := database.Query("SELECT * FROM accounts")
	// var accountId string
	// var userId string
	// var accountType string
	// var balance string
	// var status string

	// for rows.Next() {
	// 	rows.Scan(&accountId, &userId, &accountType, &balance, &status)
	// 	log.Println(accountId + " | " + userId + " | " + accountType + " | " + balance + " | " + status)
	// }
}
