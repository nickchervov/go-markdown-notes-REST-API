package adapters

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/domain"
)

type Postgres struct {
	db *sqlx.DB
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("creating migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://pkg/migrations", "pgx", driver)
	if err != nil {
		return fmt.Errorf("creating migration: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration up: %w", err)
	}
	return nil
}

func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	nativeDb := stdlib.OpenDB(*config)
	db := sqlx.NewDb(nativeDb, "postgres")

	if err := migrateUp(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("up migration: %w", err)
	}

	repo := Postgres{db: db}
	return &repo, nil
}

func (p *Postgres) CreateNote(ctx context.Context, note domain.Note) (domain.Note, error) {
	query := `INSERT INTO notes (title, content, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	if err := p.db.GetContext(ctx, &id, query, note.Title, note.Content, note.Tags, note.CreatedAt, note.UpdatedAt); err != nil {
		return domain.Note{}, fmt.Errorf("creating note: %w", err)
	}

	queryGetCreatedNote := "SELECT id, title, content, tags, created_at, updated_at FROM notes WHERE id = $1"
	var createdNote domain.Note
	if err := p.db.GetContext(ctx, &createdNote, queryGetCreatedNote, id); err != nil {
		return domain.Note{}, fmt.Errorf("getting created note: %w", err)
	}

	return createdNote, nil
}

func (p *Postgres) GetNotes(ctx context.Context, page, limit int) ([]domain.Note, int, error) {
	offset := (page - 1) * limit

	query := "SELECT id, title, content, tags, created_at, updated_at FROM notes ORDER BY id LIMIT $1 OFFSET $2"
	var notes []domain.Note
	if err := p.db.SelectContext(ctx, &notes, query, limit, offset); err != nil {
		return nil, 0, fmt.Errorf("getting notes: %w", err)
	}
	if notes == nil {
		return []domain.Note{}, 0, nil
	}

	queryTotal := "SELECT Count(*) FROM notes"
	var total int
	if err := p.db.GetContext(ctx, &total, queryTotal); err != nil {
		return nil, 0, fmt.Errorf("getting total count notes: %w", err)
	}
	return notes, total, nil
}

func (p *Postgres) GetNotesBySearch(ctx context.Context, search string) ([]domain.Note, error) {
	searchForDb := "%" + search + "%"

	query := "SELECT id, title, content, tags, created_at, updated_at FROM notes WHERE title LIKE $1"
	var notes []domain.Note
	if err := p.db.SelectContext(ctx, &notes, query, searchForDb); err != nil {
		return nil, fmt.Errorf("getting notes by searching: %w", err)
	}
	if notes == nil {
		return []domain.Note{}, nil
	}
	return notes, nil
}

func (p *Postgres) GetNote(ctx context.Context, id int) (domain.Note, error) {
	query := "SELECT id, title, content, tags, created_at, updated_at FROM notes WHERE id = $1"
	var note domain.Note
	if err := p.db.GetContext(ctx, &note, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Note{}, domain.ErrNoteNotFound
		}
		return domain.Note{}, fmt.Errorf("getting note: %w", err)
	}
	return note, nil
}

func (p *Postgres) UpdateNote(ctx context.Context, note domain.Note) error {
	query := "UPDATE notes SET title = $1, content = $2, tags = $3, updated_at = $4 WHERE id = $5"
	_, err := p.db.ExecContext(ctx, query, note.Title, note.Content, note.Tags, note.UpdatedAt, note.Id)
	if err != nil {
		return fmt.Errorf("updating note: %w", err)
	}
	return nil
}

func (p *Postgres) DeleteNote(ctx context.Context, id int) error {
	query := "DELETE FROM notes WHERE id = $1"
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Deleting note: %w", err)
	}
	return nil
}

func (p *Postgres) ExportNote(ctx context.Context, id int) (domain.Note, error) {
	query := "SELECT title, content FROM notes WHERE id = $1"
	var note domain.Note
	if err := p.db.GetContext(ctx, &note, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Note{}, domain.ErrNoteNotFound
		}
		return domain.Note{}, fmt.Errorf("getting note for export: %w", err)
	}
	return note, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
