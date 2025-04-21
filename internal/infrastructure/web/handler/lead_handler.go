package handler

import (
	"eleven-leads-api/internal/application/usecase"
	"eleven-leads-api/internal/domain/lead"
	"encoding/json"
	"net/http"
)

type LeadHandler struct {
	Usecase *usecase.LeadUsecase
}

func NewLeadHandler(u *usecase.LeadUsecase) *LeadHandler {
	return &LeadHandler{Usecase: u}
}

// CreateLead godoc
// @Summary      Cria um novo lead
// @Description  Recebe os dados de um lead e salva no sistema
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        lead body lead.Lead true "Dados do Lead"
// @Success      201 {string} string "created"
// @Failure      400 {string} string "erro de validação"
// @Router       /leads [post]
func (h *LeadHandler) CreateLead(w http.ResponseWriter, r *http.Request) {
	var l lead.Lead
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.Usecase.CapturarLead(&l); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))
}
