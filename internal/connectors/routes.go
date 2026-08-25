package connectors

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/service"
)

func SetRoutes(svc *service.NotesService) http.Handler {
	r := chi.NewRouter()

	h := New(svc)

	r.Use(middleware.Logger)
	r.Use(h.RateLimit)

	r.Post("/api/notes", h.CreateNote)
	r.Get("/api/notes", h.ListNotes)
	r.Get("/api/notes/search", h.ListNotesBySearch)

	r.Get("/api/notes/{id}", h.GetNote)
	r.Put("/api/notes/{id}", h.UpdateNote)
	r.Delete("/api/notes/{id}", h.DeleteNote)
	r.Get("/api/notes/{id}/export", h.ExportNote)

	return r
}
