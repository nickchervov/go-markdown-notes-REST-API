package service

import (
	"context"
	"fmt"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) ExportNote(ctx context.Context, input dto.ExportNoteInput) (dto.ExportNoteOutput, error) {
	if input.Id < 0 {
		return dto.ExportNoteOutput{}, domain.ErrIncorrectId
	}

	note, err := s.repo.ExportNote(ctx, input.Id)
	if err != nil {
		return dto.ExportNoteOutput{}, fmt.Errorf("export note: %w", err)
	}

	output := dto.ExportNoteOutput{
		Title:   note.Title,
		Content: note.Content,
	}

	return output, nil
}
