package models

type NoteDto struct {
	Content string `json:"conten"`
	Order   string `json:"order"`
	TaskId  string `json:"taskId"`
	Id      string `json:"id"`
}
