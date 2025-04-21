package db

import (
	"eleven-leads-api/internal/domain/lead"
	"fmt"
)

// FakeRepository implementa lead.Repository
type FakeRepository struct {
	dados []*lead.Lead
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{
		dados: make([]*lead.Lead, 0),
	}
}

func (f *FakeRepository) Salvar(l *lead.Lead) error {
	f.dados = append(f.dados, l)
	fmt.Printf("📥 Lead salvo em memória: %+v\n", l)
	return nil
}
