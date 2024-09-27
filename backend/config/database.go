package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/deathstarset/backend-docu-quest/database"
	_ "github.com/lib/pq"
	_ "github.com/pgvector/pgvector-go"
)

var DB *database.Queries
var Client *sql.DB

func StartPostgres() error {
	dbUrl := os.Getenv("DB_STRING")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil
	}

	err = db.Ping()
	log.Println("Connected to database succefully")
	if err != nil {
		return err
	}

	queries := database.New(db)

	DB = queries
	Client = db

	return nil
}

func ClosePostgres() error {
	return nil
}
