package store

import (
	"github.com/TretyakovArtem/lms/internal/app/model"
)

// ReportRepository ...
type ReportRepository struct {
	store *Store
}

// Rarely ..
func (r *ReportRepository) Rarely(dateFrom string, dateTo string) ([]model.Book, error) {
	rows, err := r.store.db.Query("SELECT DISTINCT name, author, code FROM books JOIN log ON books.id = log.book_id WHERE books.id NOT IN (SELECT log.id FROM log WHERE date_from > ? AND date_to < ?)", dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	var books []model.Book

	for rows.Next() {
		var book model.Book
		_ = rows.Scan(&book.Name, &book.Author, &book.Code)
		books = append(books, book)
	}

	return books, nil
}

// Rate ..
func (r *ReportRepository) Rate() ([]model.Book, error) {
	rows, err := r.store.db.Query("SELECT books.id, name, author, COUNT(*) as rate FROM books JOIN log ON books.id = log.book_id GROUP BY books.id ORDER BY rate DESC")
	if err != nil {
		return nil, err
	}

	var books []model.Book

	for rows.Next() {
		var book model.Book
		_ = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Rate)
		books = append(books, book)
	}

	return books, nil
}

// Credit ..
func (r *ReportRepository) Credit() (interface{}, error) {

	rows, err := r.store.db.Query("SELECT log.id, customers.name, books.name FROM log JOIN customers ON log.customer_id = customers.id JOIN books ON log.book_id = books.id WHERE on_store = 0 AND date_to < NOW();")
	if err != nil {
		return nil, err
	}

	type creditCard struct {
		ID           int
		BookName     string
		CustomerName string
	}

	var cards []creditCard

	for rows.Next() {
		var c creditCard
		_ = rows.Scan(&c.ID, &c.CustomerName, &c.BookName)
		cards = append(cards, c)
	}

	return cards, nil
}
