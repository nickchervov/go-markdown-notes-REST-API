package service

import (
	"context"
	"fmt"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) DeleteNote(ctx context.Context, input dto.DeleteNoteInput) error {
	if input.Id < 0 {
		return domain.ErrIncorrectId
	}

	if err := s.repo.DeleteNote(ctx, input.Id); err != nil {
		return fmt.Errorf("deleting note: %w", err)
	}

	return nil
}
