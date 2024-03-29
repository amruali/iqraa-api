package handlers

import (
	"iqraa-api/domain"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	domain *domain.Domain
}

func NewServer(domain *domain.Domain) *Server {
	return &Server{domain: domain}
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(120 * time.Second))
}

func SetupRouter(domain *domain.Domain) *chi.Mux {
	server := NewServer(domain)
	r := chi.NewRouter()

	setupMiddleware(r)
	server.setupEndpoints(r)

	return r
}
