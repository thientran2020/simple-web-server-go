package database

import (
	"database/sql"
	"log"
)

func Connect() {
	_, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic("Could not connect to database.db!")
	}
	log.Println("Database connected...!!!")

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
