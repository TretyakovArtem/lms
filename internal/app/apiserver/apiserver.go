package apiserver

import (
	"html/template"
	"net/http"
	"os"

	"github.com/TretyakovArtem/lms/internal/app/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {

	bh := NewBookHandler(s.store)

	s.router.HandleFunc("/books", bh.Index())


}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) index() http.HandlerFunc {
	wd, _ := os.Getwd()
	tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/index.gohtml"))

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", 3)
	}
}

func (s *APIServer) reports() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		wd, _ := os.Getwd()
		tpl := template.Must(template.ParseFiles(wd + "/internal/app/view/index_reports.gohtml"))
		tpl.ExecuteTemplate(w, "index_reports.gohtml", 3)
	}
}
