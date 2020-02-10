package db

import (
	"database/sql"
	"github.com/ebladrocher/smarthome/system/store"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:db,
	}
}

// ConnectToDB ...
func ConnectToDB(db *ConfigDB) (*sql.DB, error) {

	tmp := func (dbURL string) (*sql.DB, error) {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			return nil, err
		}

		if err := db.Ping(); err != nil {
			return nil, err
		}

		return db, nil
	}

 	qwerty, err := tmp(db.ConnectionSting())
 	if err != nil {
 		return nil, err
	}

 	return qwerty, nil

}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}


