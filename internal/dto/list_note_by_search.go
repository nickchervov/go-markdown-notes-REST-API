package dto

import "github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"

type ListNotesBySearchInput struct {
	Search string
}

type ListNotesBySearchOutput struct {
	Notes []domain.Note `json:"notes"`
}
