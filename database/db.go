package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Change credentials for your DB
	connStr := "host=localhost user= password= dbname=gin_auth port=5432 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach DB:", err)
	}

	fmt.Println("Connected to database!")

}
