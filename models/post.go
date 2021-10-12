package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	PostId    uint   `json:"postId" gorm:"autoIncrement"`
	CreatedBY uint   `json:"user_id"`
	Content   string `json:"content"`
}

type CreatePostInput struct {
	CreatedBy uint   `json:"user_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
