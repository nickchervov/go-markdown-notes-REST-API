package service

import (
	"context"
	"fmt"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) GetNote(ctx context.Context, input dto.GetNoteInput) (dto.GetNoteOutput, error) {
	if input.Id <= 0 {
		return dto.GetNoteOutput{}, domain.ErrIncorrectId
	}

	note, err := s.repo.GetNote(ctx, input.Id)
	if err != nil {
		return dto.GetNoteOutput{}, fmt.Errorf("get note: %w", err)
	}

	output := dto.GetNoteOutput{
		Id:        note.Id,
		Title:     note.Title,
		Content:   note.Content,
		Tags:      note.Tags,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	return output, nil
}
