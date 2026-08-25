package service

import (
	"context"

	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
)

type PostgresRepository interface {
	CreateNote(ctx context.Context, note domain.Note) (domain.Note, error)
	GetNotes(ctx context.Context, page, limit int) ([]domain.Note, int, error)
	GetNotesBySearch(ctx context.Context, search string) ([]domain.Note, error)
	GetNote(ctx context.Context, id int) (domain.Note, error)
	UpdateNote(ctx context.Context, note domain.Note) error
	DeleteNote(ctx context.Context, id int) error
	ExportNote(ctx context.Context, id int) (domain.Note, error)
}

type RedisCache interface {
	RateLimit(ctx context.Context, rpm int, ip string) (int, bool, error)
}

type NotesService struct {
	repo  PostgresRepository
	cache RedisCache
}

func New(repo PostgresRepository, cache RedisCache) *NotesService {
	return &NotesService{repo: repo, cache: cache}
}
