package apiserver

import (
	"net/http"

	"github.com/TretyakovArtem/lms/internal/app/store"
)

// BookHandler ...
type BookHandler struct {
	store *store.Store
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
		//
	}
}
