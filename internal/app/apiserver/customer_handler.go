package apiserver

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/TretyakovArtem/lms/internal/app/model"
	"github.com/TretyakovArtem/lms/internal/app/store"
)

// CustomerHandler ...
type CustomerHandler struct {
	store *store.Store
}

// NewCustomerHandler создает новый обьект.
func NewCustomerHandler(store *store.Store) *CustomerHandler {
	return &CustomerHandler{
		store: store,
	}
}

// Index ..
func (h *CustomerHandler) Index() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/customer/customer.gohtml"))

		customers, _ := h.store.Customer().List()

		tpl.ExecuteTemplate(w, "customer.gohtml", customers)
	}
}

// Create ...
func (h *CustomerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/customer/create_customer.gohtml"))

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			phone := r.FormValue("phone")
			address := r.FormValue("address")

			customer := &model.Customer{
				ID:      1,
				Name:    name,
				Phone:   phone,
				Address: address,
			}

			h.store.Customer().Create(customer)

			http.Redirect(w, r, "/customers", http.StatusSeeOther)
		}

		tpl.ExecuteTemplate(w, "create_customer.gohtml", 3)
	}
}

// Update ...
func (h *CustomerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/customer/update_customer.gohtml"))

		customerID := r.URL.Query().Get("id")

		cID, _ := strconv.Atoi(customerID)

		customer, _ := h.store.Customer().FindByID(customerID)

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			phone := r.FormValue("phone")
			address := r.FormValue("address")

			customer := &model.Customer{
				Name:    name,
				Phone:   phone,
				Address: address,
				ID:      cID,
			}

			h.store.Customer().Update(customer)

			http.Redirect(w, r, "/customers", http.StatusSeeOther)

		}
		tpl.ExecuteTemplate(w, "update_customer.gohtml", customer)
	}
}

// Delete ...
func (h *CustomerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		customerID := r.URL.Query().Get("id")

		cID, _ := strconv.Atoi(customerID)

		h.store.Customer().Delete(cID)

		http.Redirect(w, r, "/customers", http.StatusSeeOther)

	}

}
