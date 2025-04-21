package usecase

import (
	"eleven-leads-api/internal/domain/lead"
	"errors"
	"regexp"
	"time"
)

type LeadUsecase struct {
	Repo lead.Repository
}

func NewLeadUsecase(repo lead.Repository) *LeadUsecase {
	return &LeadUsecase{Repo: repo}
}

func (u *LeadUsecase) CapturarLead(l *lead.Lead) error {
	if l.Nome == "" || l.Email == "" {
		return errors.New("nome e email são obrigatórios")
	}
	if !isValidEmail(l.Email) {
		return errors.New("email inválido")
	}

	l.CriadoEm = time.Now()

	return u.Repo.Salvar(l)
}

func isValidEmail(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(email)
}
