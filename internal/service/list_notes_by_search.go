package service

import (
	"context"
	"fmt"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) ListNotesBySearch(ctx context.Context, input dto.ListNotesBySearchInput) (dto.ListNotesBySearchOutput, error) {
	notes, err := s.repo.GetNotesBySearch(ctx, input.Search)
	if err != nil {
		return dto.ListNotesBySearchOutput{}, fmt.Errorf("getting notes by searching: %w", err)
	}

	output := dto.ListNotesBySearchOutput{
		Notes: notes,
	}

	return output, nil
}
