package postgrestore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// InitDB ...
func InitDB(dbUrl string) *sql.DB {
	db,err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatalf("Error occured while establishing connection to db")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error")
	}

	return db

}


