package web

import (
	"eleven-leads-api/internal/infrastructure/web/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(leadHandler *handler.LeadHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/api/leads", leadHandler.CreateLead)

	return r
}
