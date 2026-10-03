package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/render"
)

func (h *NotesHandler) ListNotes(w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")
	limit := r.FormValue("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "100"
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		render.JSON(w, domain.ErrIncorrectPageOrLimit.Code, domain.ErrIncorrectPageOrLimit)
		return
	}
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		render.JSON(w, domain.ErrIncorrectPageOrLimit.Code, domain.ErrIncorrectPageOrLimit)
		return
	}

	input := dto.ListNotesInput{
		Page:  pageNum,
		Limit: limitNum,
	}

	output, err := h.svc.ListNotes(r.Context(), input)
	if err != nil {
		var targetErr *domain.NoteError
		if errors.As(err, &targetErr) {
			render.JSON(w, targetErr.Code, targetErr)
			return
		}
		log.Println(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"message": "get notes internal error"})
		return
	}

	render.JSON(w, http.StatusOK, output)
}
