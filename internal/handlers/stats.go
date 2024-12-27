package handlers

import (
	"net/http"

	"github.com/hydrogen-1/goth-stats/internal/repository"
	"github.com/hydrogen-1/goth-stats/internal/templates"
)

type StatsHandler struct {
	Repository *repository.Queries
}

func (h *StatsHandler) Visit(w http.ResponseWriter, r *http.Request) {
	err := h.Repository.InsertVisit(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	total, err := h.Repository.GetTotal(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	last24h, err := h.Repository.GetLast24h(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	templates.RenderStats(w, total, last24h)
}
