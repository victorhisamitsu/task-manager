package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"tasks"`
	ID            string     `json:"ID" bun:"id,pk,type:uuid"`
	Title         string     `json:"title" bun:"Title"`
	Status        string     `json:"status" bun:"Status"`
	DueDate       *time.Time `json:"due_date" bun:"duedate"`
	Importante    bool       `json:"importante" bun:"Importante"`
	CreatedAt     time.Time  `json:"created_at" bun:"CreatedAt"`
	Notes         []Note     `json:"notes" bun:"Notes"`
}
