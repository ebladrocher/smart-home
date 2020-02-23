package postgrestore

import (
	"database/sql"
	"github.com/ebladrocher/smarthome/system/config"
	_ "github.com/lib/pq"
	"log"
)

// InitDB ...
func InitDB() *sql.DB {

	//tmp := func (dbURL string) (*sql.DB, error) {
	//	db, err := sql.Open("postgres", dbURL)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if err := db.Ping(); err != nil {
	//		return nil, err
	//	}
	//
	//	return db, nil
	//}
	//
 	//str, err := tmp(db.ConnectionSting())
 	//if err != nil {
 	//	return nil, err
	//}
	//
 	//return str, nil

	cfg, _ := config.ReadConfig()
	dbUrl := NewDbConfig(cfg)
	db,err := sql.Open("postgres", dbUrl.ConnectionSting())

	if err != nil {
		log.Fatalf("Error occured while establishing connection to db")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error")
	}

	return db

}


