package handlers

import (
	"github.com/HHHMHA/celeritas"
	"net/http"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
