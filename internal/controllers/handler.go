package connectors

import "github.com/nickchervov/go-markdown-notes-REST-API/internal/service"

type NotesHandler struct {
	svc *service.NotesService
}

func New(svc *service.NotesService) *NotesHandler {
	return &NotesHandler{svc: svc}
}
