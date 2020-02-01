package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db *sql.DB
}

func Init() *Store {
	return &Store{}
}

func New() *Store {
	tmp, _ := ReadConfig()
	var s = Store{config: tmp}
	return &s
}


func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabasseUrl)
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
	st := New()
	if err := st.Open(); err != nil {
		return err
	}

	s.db=st.db
	s.config = st.config
	return nil
}


