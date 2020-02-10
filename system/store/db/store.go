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

//func (s *Store) Open() error {
//  db, err := sql.Open("postgres", s.config.ConnectionSting())
//  if err != nil {
//    return err
//  }
//  if err := db.Ping(); err != nil {
//    return err
//  }
//  s.db = db
//  return nil
//}

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


