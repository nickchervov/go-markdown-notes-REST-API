package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, domain.ErrIncorrectId.Code, domain.ErrIncorrectId)
		return
	}

	input := dto.DeleteNoteInput{
		Id: id,
	}

	if err := h.svc.DeleteNote(r.Context(), input); err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		log.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "delete note internal error"})
		return
	}

	render.JSON(w, http.StatusNoContent, nil)
}
