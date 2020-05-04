package store

import "github.com/p4thf1nderr/lms/internal/app/model"

// BookRepository ...
type BookRepository interface {
	Create(*model.Book) error
	Index() ([]model.Book, error)
}
