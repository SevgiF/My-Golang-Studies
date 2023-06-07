package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	username := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, port, username, password, dbName, sslmode)

	conn, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func CloseDatabase(conn *sql.DB) {
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
