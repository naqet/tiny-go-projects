package handlers

import (
	"naqet/forms/views/home"
	"net/http"
)

type PagesHandler struct{}

func NewPagesHandler() *PagesHandler {
	return &PagesHandler{}
}

func (h *PagesHandler) Home(w http.ResponseWriter, r *http.Request) {
    home.Index().Render(r.Context(), w)
}
