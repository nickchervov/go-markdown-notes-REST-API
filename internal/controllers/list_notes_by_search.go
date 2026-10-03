package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) ListNotesBySearch(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("q")

	input := dto.ListNotesBySearchInput{Search: search}

	output, err := h.svc.ListNotesBySearch(r.Context(), input)
	if err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		log.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "get notes by search internal error"})
		return
	}
	render.JSON(w, http.StatusOK, output)
}
