package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) UpdateNote(ctx context.Context, input dto.UpdateNoteInput) error {
	if input.Id <= 0 {
		return domain.ErrIncorrectId
	}
	note, err := domain.NewNote(input.Title, input.Content, input.Tags)
	if err != nil {
		return err
	}
	note.Id = input.Id
	note.UpdatedAt = time.Now()

	if err := s.repo.UpdateNote(ctx, note); err != nil {
		return fmt.Errorf("updating note: %w", err)
	}
	return nil
}
