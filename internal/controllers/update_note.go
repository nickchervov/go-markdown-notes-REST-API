package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, domain.ErrIncorrectId.Code, domain.ErrIncorrectId)
		return
	}

	var input dto.UpdateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"message": "decoding request body: " + err.Error()})
		return
	}
	input.Id = id

	if err := h.svc.UpdateNote(r.Context(), input); err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		log.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "update note internal error"})
		return
	}
	render.JSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("id %d updated", id)})
}
