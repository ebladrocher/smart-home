package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config         *DbConfig
	db             *sql.DB
	userRepository *UserRepository
}

func InitStore(cfg *DbConfig) *Store {
	return &Store{config: cfg}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.ConnectionSting())
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

func (s *Store) SetStore() error {
	if err := s.Open(); err != nil {
		return err
	}

	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository !=  nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store:s,
	}

	return s.userRepository
}


