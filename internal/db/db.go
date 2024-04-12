package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(connectionurl string) *sql.DB {
	db, err := sql.Open("postgres", connectionurl)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	return db
}

func CreateMoneyTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS money (
        id VARCHAR PRIMARY KEY,
        amount FLOAT NOT NULL DEFAULT 0
    );`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Error creating money table: %v", err)
	}
}
