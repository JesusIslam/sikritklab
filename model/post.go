package model

import (
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/database"
)

func init() {
	err := database.DB.Init(&Post{})
	if err != nil {
		lowger.Fatal("Failed to initialize database post:", err)
	}
}

type Post struct {
	ID        int       `json:"id" storm:"id,increment"`
	CreatedAt time.Time `json:"created_at" storm:"index"`
	ThreadID  string    `json:"thread_id" storm:"index"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content"`
}
