package domain

type NoteError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e NoteError) Error() string {
	return e.Message
}

var (
	ErrValidation           = &NoteError{Code: 400, Message: "invalid validation"}
	ErrIncorrectPageOrLimit = &NoteError{Code: 400, Message: "incorrect page or limit"}
	ErrIncorrectId          = &NoteError{Code: 400, Message: "incorrect id"}
	ErrNoteNotFound         = &NoteError{Code: 404, Message: "note not found"}
)
