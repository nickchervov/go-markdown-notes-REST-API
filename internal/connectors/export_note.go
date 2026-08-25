package connectors

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) ExportNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, domain.ErrIncorrectId.Code, domain.ErrIncorrectId)
		return
	}

	input := dto.ExportNoteInput{
		Id: id,
	}

	output, err := h.svc.ExportNote(r.Context(), input)
	if err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "export note internal error"})
		return
	}

	fileContent := []byte(output.Content)
	fileReader := bytes.NewReader(fileContent)

	filename := fmt.Sprintf("%s.md", output.Title)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	http.ServeContent(w, r, filename, time.Now(), fileReader)
}
