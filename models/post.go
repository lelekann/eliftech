package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	CreatedBY uint   `json:"user_id"`
	Content   string `json:"content"`
	User      User   `json:"ID" gorm:"foreignKey:CreatedBy"`
}

type CreatePostInput struct {
	CreatedBy uint   `json:"user_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
