package store

import (
	"math/rand"

	"github.com/TretyakovArtem/lms/internal/app/model"
)

// BookRepository ...
type BookRepository struct {
	store *Store
}

// Create метод для занесения в базу читателя
func (r *BookRepository) Create(c *model.Book) (*model.Book, error) {

	min := 10
	max := 30
	code := rand.Intn(max-min) + min

	r.store.db.QueryRow("INSERT INTO books (name, author, code, year, publisher, count_papers, total_balance) VALUES (?, ?, ?, ?, ?, ?, ?)",
		c.Name, c.Author, code, c.Year, c.Publisher, c.CountPapers, c.TotalBalance)
	return nil, nil
}

// List список читателей
func (r *BookRepository) List() ([]model.Book, error) {

	rows, err := r.store.db.Query(`SELECT id, name, author, code, year, publisher, count_papers, total_balance FROM books`)
	if err != nil {
		return nil, err
	}

	var books []model.Book

	for rows.Next() {
		var book model.Book
		_ = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Code, &book.Year, &book.Publisher, &book.CountPapers, &book.TotalBalance)
		books = append(books, book)
	}

	return books, nil
}

// FindByID ..
func (r *BookRepository) FindByID(id string) (*model.Book, error) {
	b := &model.Book{}

	if err := r.store.db.QueryRow("SELECT id, name, author, code, year, publisher, count_papers, total_balance FROM books WHERE id = ?", id).Scan(
		&b.ID,
		&b.Name,
		&b.Author,
		&b.Code,
		&b.Year,
		&b.Publisher,
		&b.CountPapers,
		&b.TotalBalance,
	); err != nil {
		return nil, err
	}

	return b, nil
}

// Update метод для изменения данных в базе читателя
func (r *BookRepository) Update(b *model.Book) (*model.Book, error) {
	r.store.db.QueryRow("UPDATE books SET name=?, author=?, code=?, year=?, publisher=?, count_papers=?, total_balance=? WHERE id=?",
		b.Name,
		b.Author,
		b.Code,
		b.Year,
		b.Publisher,
		b.CountPapers,
		b.TotalBalance,
		b.ID,
	)
	return nil, nil
}

// Delete метод для удаления книги
func (r *BookRepository) Delete(bID int) (*model.Book, error) {
	r.store.db.QueryRow("DELETE FROM books WHERE id = ?", bID)
	return nil, nil
}

// SetCustomer метод для изменения данных в базе читателя
func (r *BookRepository) SetCustomer(bookID int, customerID int, dateFrom string, dateTo string, onStore int) (bool, error) {
	r.store.db.QueryRow("INSERT INTO log (book_id, customer_id, date_from, date_to, on_store) VALUES (?, ?, ?, ?, ?)",
		bookID, customerID, dateFrom, dateTo, onStore)
	return true, nil
}
