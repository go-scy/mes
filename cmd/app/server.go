package app

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	csrfKey = "csrf-auth-key"
)

type server struct {
	router *mux.Router
	config *AppConfig
}

func NewServer(config *AppConfig) *server {
	return &server{
		mux.NewRouter(),
		config,
	}
}

func (s server) AddMiddleWares() {
	s.router.Use(middleWare)
}

func (s server) AddCsrfMiddleware() {
	csrfMiddleWare := csrf.Protect([]byte(csrfKey))
	s.router.Use(csrfMiddleWare)
}

func (s server) AddRoutes() {
	s.router.Handle("/", rootHandler(s.config))
	s.router.Handle("/employee", employeeHandler())
	s.router.Handle("/product", productHandler())
}

func (s server) Router() *mux.Router {
	return s.router
}
