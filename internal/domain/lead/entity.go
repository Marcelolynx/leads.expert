package lead

import "time"

type Lead struct {
	ID       int64     `json:"id"`
	Nome     string    `json:"nome" example:"Marcelo Oliveira"`
	Email    string    `json:"email" example:"marcelo@email.com"`
	Origem   string    `json:"origem" example:"LandingPage"`
	CriadoEm time.Time `json:"criado_em"`
}
