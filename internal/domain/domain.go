package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Note struct {
	Id        int       `db:"id"`
	Title     string    `db:"title" validate:"required,max=128"`
	Content   string    `db:"content"`
	Tags      string    `db:"tags" validate:"required"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var validate = validator.New()

func NewNote(title, content, tags string) (Note, error) {
	note := Note{
		Title:   title,
		Content: content,
		Tags:    tags,
	}

	if err := validate.Struct(note); err != nil {
		return Note{}, ErrValidation
	}

	return note, nil
}
