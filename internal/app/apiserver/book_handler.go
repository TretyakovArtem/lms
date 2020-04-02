package apiserver

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/TretyakovArtem/lms/internal/app/model"
	"github.com/TretyakovArtem/lms/internal/app/store"
)

// BookHandler ...
type BookHandler struct {
	store *store.Store
}

// Context ...
type Context struct {
	Book      *model.Book
	Customers []model.Customer
	DateFrom  string
	DateTo    string
	OnStore   int
}

// NewBookHandler создает новый обьект.
func NewBookHandler(store *store.Store) *BookHandler {
	return &BookHandler{
		store: store,
	}
}

// Index ..
func (h *BookHandler) Index() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/book/book.gohtml"))

		books, _ := h.store.Book().List()

		tpl.ExecuteTemplate(w, "book.gohtml", books)
	}
}

// Create ...
func (h *BookHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/book/create_book.gohtml"))

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			author := r.FormValue("author")
			code := r.FormValue("code")
			year := r.FormValue("year")
			publisher := r.FormValue("publisher")

			countPapers := r.FormValue("count_papers")
			totalBalance := r.FormValue("total_balance")

			cp, _ := strconv.Atoi(countPapers)
			tb, _ := strconv.Atoi(totalBalance)

			book := &model.Book{
				ID:           1,
				Name:         name,
				Author:       author,
				Code:         code,
				Year:         year,
				Publisher:    publisher,
				CountPapers:  cp,
				TotalBalance: tb,
			}

			h.store.Book().Create(book)

			http.Redirect(w, r, "/books", http.StatusSeeOther)
		}

		tpl.ExecuteTemplate(w, "create_book.gohtml", 3)
	}
}

// Update ...
func (h *BookHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/book/update_book.gohtml"))

		bookID := r.URL.Query().Get("id")

		bID, _ := strconv.Atoi(bookID)

		book, _ := h.store.Book().FindByID(bookID)

		customers, _ := h.store.Customer().List()

		var dateFrom, dateTo string
		var onS int

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			author := r.FormValue("author")
			code := r.FormValue("code")
			year := r.FormValue("year")
			publisher := r.FormValue("publisher")

			customerID := r.FormValue("customer")

			countPapers := r.FormValue("count_papers")
			totalBalance := r.FormValue("total_balance")
			dateFrom := r.FormValue("date_from")
			dateTo := r.FormValue("date_to")
			onStore := r.FormValue("on_store")

			cp, _ := strconv.Atoi(countPapers)
			tb, _ := strconv.Atoi(totalBalance)
			onS, _ := strconv.Atoi(onStore)

			cID, _ := strconv.Atoi(customerID)

			book := &model.Book{
				ID:           bID,
				Name:         name,
				Author:       author,
				Code:         code,
				Year:         year,
				Publisher:    publisher,
				CountPapers:  cp,
				TotalBalance: tb,
			}

			h.store.Book().Update(book)

			h.store.Book().SetCustomer(bID, cID, dateFrom, dateTo, onS)

			// запись в лог
			http.Redirect(w, r, "/books", http.StatusSeeOther)

		}
		ctx := &Context{
			Book:      book,
			Customers: customers,
			DateFrom:  dateFrom,
			DateTo:    dateTo,
			OnStore:   onS,
		}
		tpl.ExecuteTemplate(w, "update_book.gohtml", ctx)
	}
}

// Delete ...
func (h *BookHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bookID := r.URL.Query().Get("id")

		bID, _ := strconv.Atoi(bookID)

		h.store.Book().Delete(bID)

		http.Redirect(w, r, "/books", http.StatusSeeOther)

	}

}
