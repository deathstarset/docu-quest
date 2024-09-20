package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/deathstarset/backend-docu-quest/database"
	_ "github.com/lib/pq"
)

var DB *database.Queries

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

	return nil
}

func ClosePostgres() error {
	return nil
}
