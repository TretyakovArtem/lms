package sqlstore

import (
	"database/sql"

	"github.com/p4thf1nderr/lms/internal/app/store"
	_ "github.com/lib/pq"
)

// Store ..
type Store struct {
	db             *sql.DB
	bookRepository *BookRepository
}

// New ..
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Book ...
func (s *Store) Book() store.BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{
		store: s,
	}

	return s.bookRepository
}
