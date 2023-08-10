package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Note struct {
	bun.BaseModel `bun:"notes"`
	ID            string    `json:"ID" bun:"id,pk,type:uuid"`
	Content       string    `json:"content" bun:"content"`
	CreatedAt     time.Time `json:"created_at" bun:"CreatedAt"`
	Order         string    `json:"order" bun:"Order"`
	TaskId        string    `json:"task_id" bun:"taskid"`
}
