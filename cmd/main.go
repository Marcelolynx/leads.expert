// @title Eleven Leads API
// @version 1.0
// @description API para captação de leads do projeto ELEVEN LEADS
// @termsOfService http://elevenleads.com.br/terms

// @contact.name Suporte Eleven Experts
// @contact.email suporte@elevenexperts.com.br

// @host localhost:8080
// @BasePath /api

package main

import (
	"log"
	"net/http"

	_ "eleven-leads-api/docs"

	"eleven-leads-api/internal/application/usecase"
	"eleven-leads-api/internal/infrastructure/db" // será usado logo
	"eleven-leads-api/internal/infrastructure/web"
	"eleven-leads-api/internal/infrastructure/web/handler"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := chi.NewRouter()

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Ping
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// Setup Lead
	repo := db.NewFakeRepository() // vamos substituir por Postgres depois
	uc := usecase.NewLeadUsecase(repo)
	leadHandler := handler.NewLeadHandler(uc)

	r.Mount("/", web.NewRouter(leadHandler))

	log.Println("🚀 Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)
}
