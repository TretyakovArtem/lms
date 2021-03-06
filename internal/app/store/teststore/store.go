package teststore

import (
	"github.com/p4thf1nderr/lms/internal/app/model"
	"github.com/p4thf1nderr/lms/internal/app/store"
)

// Store ..
type Store struct {
	bookRepository *BookRepository
}

// New ..
func New() *Store {
	return &Store{}
}

// Book ...
func (s *Store) Book() store.BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{
		store: s,
		books: make(map[string]*model.Book),
	}

	return s.bookRepository
}
