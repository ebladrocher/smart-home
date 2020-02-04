package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config *DbConfig
	db *sql.DB
}

func Init() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	tmp := NewDbConfig()
	db, err := sql.Open("postgres", tmp.ConnectionSting())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) ConfigureStore() error {
	if err := s.Open(); err != nil {
		return err
	}

	return nil
}


