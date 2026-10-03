package connectors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"message": "decoding request body: " + err.Error()})
		return
	}

	output, err := h.svc.CreateNote(r.Context(), input)
	if err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		log.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "create note internal error"})
		return
	}

	render.JSON(w, http.StatusCreated, output)
}
