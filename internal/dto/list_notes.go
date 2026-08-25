package dto

import "github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"

type ListNotesInput struct {
	Page  int
	Limit int
}

type ListNotesOutput struct {
	Notes []domain.Note `json:"notes"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}
