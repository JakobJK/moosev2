package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error validating database arguments: %v", err)
	}

	// --- CONNECTION POOL CONFIGURATION ---
	
	// MaxOpenConns: Maximum number of established connections to the database.
	DB.SetMaxOpenConns(25)

	// MaxIdleConns: Maximum number of connections in the idle connection pool.
	DB.SetMaxIdleConns(10)

	// ConnMaxLifetime: Maximum amount of time a connection may be reused.
	DB.SetConnMaxLifetime(time.Hour)

	// ConnMaxIdleTime: Maximum amount of time a connection may be idle.
	DB.SetConnMaxIdleTime(15 * time.Minute)

	// Verify the connection is actually alive
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println(">>> DATABASE: Connected and Pool Initialized")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
