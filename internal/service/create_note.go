package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/dto"
)

func (s *NotesService) CreateNote(ctx context.Context, input dto.CreateNoteInput) (dto.CreateNoteOutput, error) {
	note, err := domain.NewNote(input.Title, input.Content, input.Tags)
	if err != nil {
		return dto.CreateNoteOutput{}, err
	}
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	createdNote, err := s.repo.CreateNote(ctx, note)
	if err != nil {
		return dto.CreateNoteOutput{}, fmt.Errorf("creating note: %w", err)
	}

	output := dto.CreateNoteOutput{
		Id:        createdNote.Id,
		Title:     createdNote.Title,
		Content:   createdNote.Content,
		Tags:      createdNote.Tags,
		CreatedAt: createdNote.CreatedAt,
		UpdatedAt: createdNote.UpdatedAt,
	}

	return output, nil
}
