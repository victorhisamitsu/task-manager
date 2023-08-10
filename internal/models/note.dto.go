package models

type NoteDto struct {
	Content string `json:"content"`
	Order   string `json:"order"`
	Id      string `json:"ID"`
}
