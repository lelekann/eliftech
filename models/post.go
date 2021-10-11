package models

import "time"

type Post struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CreatedBY uint   `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt time.Time
}

type CreatePostInput struct {
	Content string `json:"content" binding:"required"`
}
