package service

import (
	"context"
	"fmt"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) ListNotes(ctx context.Context, input dto.ListNotesInput) (dto.ListNotesOutput, error) {
	if input.Limit < 0 || input.Page < 0 {
		return dto.ListNotesOutput{}, domain.ErrIncorrectPageOrLimit
	}

	notes, total, err := s.repo.GetNotes(ctx, input.Page, input.Limit)
	if err != nil {
		return dto.ListNotesOutput{}, fmt.Errorf("get notes: %w", err)
	}

	output := dto.ListNotesOutput{
		Notes: notes,
		Total: total,
		Page:  input.Page,
		Limit: input.Limit,
	}

	return output, nil
}
