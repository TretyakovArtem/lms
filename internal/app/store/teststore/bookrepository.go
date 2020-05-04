package teststore

import "github.com/p4thf1nderr/lms/internal/app/model"

// BookRepository ...
type BookRepository struct {
	store *Store
	books map[string]*model.Book
}

// Create ...
func (r *BookRepository) Create(b *model.Book) error {
	if err := b.Validate(); err != nil {
		return err
	}

	r.books[b.Title] = b

	b.ID = uint64(len(r.books))

	return nil
}

// Index ...
func (r *BookRepository) Index() ([]model.Book, error) {

	books := []model.Book{}
	return books, nil
}
