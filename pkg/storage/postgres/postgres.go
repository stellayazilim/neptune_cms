package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type postgres struct {
	DB *sqlx.DB
}

func Postgres() *postgres {
	connString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
	)

	conn, err := sqlx.Connect("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	return &postgres{
		DB: conn,
	}
}