package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectionDB() *sql.DB {
	connection := "user=maxter dbname=store password=admin host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
