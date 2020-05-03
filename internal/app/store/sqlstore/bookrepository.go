package sqlstore

import (
	"github.com/TretyakovArtem/lms/internal/app/model"
)

// BookRepository ...
type BookRepository struct {
	store *Store
}

// Index метод для получения списка книг
func (r *BookRepository) Index() ([]model.Book, error) {
	var books []model.Book

	rows, err := r.store.db.Query("SELECT id, title, year FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Book
		rows.Scan(&book.ID, &book.Title, &book.Year)
		books = append(books, book)
	}

	return books, nil
}

// Create создание экземпляря книги
func (r *BookRepository) Create(b *model.Book) error {
	if err := b.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO books (title, year) VALUES ($1, $2) RETURNING id",
		b.Title,
		b.Year,
	).Scan(&b.ID)
}
