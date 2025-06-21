package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// 🧪 You can later use os.Getenv or a `.env` loader
	host := "localhost"
	port := 5432
	user := "gochat"
	password := "gochatpass"
	dbname := "gochat_db"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return db
}
