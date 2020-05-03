package store

import "github.com/TretyakovArtem/lms/internal/app/model"

// BookRepository ...
type BookRepository interface {
	Create(*model.Book) error
	Index() ([]model.Book, error)
}
