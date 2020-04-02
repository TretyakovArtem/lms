package apiserver

import (
	"html/template"
	"net/http"
	"os"

	"github.com/TretyakovArtem/lms/internal/app/store"
)

// ReportHandler ...
type ReportHandler struct {
	store *store.Store
}

// NewReportHandler создает новый обьект.
func NewReportHandler(store *store.Store) *ReportHandler {
	return &ReportHandler{
		store: store,
	}
}

// Rarely ..
func (h *ReportHandler) Rarely() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/reports/rarely.gohtml"))

		dateFrom := "2019-10-10"
		dateTo := "2020-01-26"

		rarely, _ := h.store.Report().Rarely(dateFrom, dateTo)

		tpl.ExecuteTemplate(w, "rarely.gohtml", rarely)
	}
}

// Rate ..
func (h *ReportHandler) Rate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/reports/rate.gohtml"))

		rate, _ := h.store.Report().Rate()

		tpl.ExecuteTemplate(w, "rate.gohtml", rate)
	}
}

// Credit ..
func (h *ReportHandler) Credit() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/reports/credit.gohtml"))

		credit, _ := h.store.Report().Credit()

		tpl.ExecuteTemplate(w, "credit.gohtml", credit)
	}
}
