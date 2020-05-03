package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ..
type Store struct {
	config         *Config
	db             *sql.DB
	bookRepository *BookRepository
}

// New ..
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ..
func (s *Store) Close() {
	s.db.Close()
}

// Book ...
func (s *Store) Book() *BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{
		store: s,
	}

	return s.bookRepository
}
