package store

import (
	"github.com/TretyakovArtem/lms/internal/app/model"
)

// BookRepository ...
type BookRepository struct {
	store *Store
}

// Index метод для получения списка книг
func (r *BookRepository) Index(c *model.Book) (*model.Book, error) {
	return nil, nil
}
