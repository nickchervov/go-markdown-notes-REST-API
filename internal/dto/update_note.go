package dto

type UpdateNoteInput struct {
	Id      int
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}
