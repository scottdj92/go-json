package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", "go:golang@tcp(127.0.0.1:11954")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
