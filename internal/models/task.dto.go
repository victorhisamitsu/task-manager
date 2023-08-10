package models

import "time"

type TaskDto struct {
	ID          string     `json:"ID"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Importante  bool       `json:"importante"`
	CreatedAt   time.Time  `json:"created_at"`
	Notes       []Note     `json:"notes"`
	Filter      string     `json:"filter_type"`
}
