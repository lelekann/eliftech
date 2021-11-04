package repositories

import (
	"eliftech_1/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id uint) (*models.Post, error)
	CreatePost(post *models.Post) (*models.Post, error)
	DeletePost(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func (r *postRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	res := r.db.Find(&posts)
	return posts, res.Error
}

func (r *postRepository) GetPostById(id uint) (*models.Post, error) {
	var post *models.Post
	res := r.db.First(&post, id)
	return post, res.Error
}

func (r *postRepository) CreatePost(post *models.Post) (*models.Post, error) {
	res := r.db.Create(&post)
	return post, res.Error
}

func (r *postRepository) DeletePost(id uint) error {
	res := r.db.Where(id).Delete(&models.Post{})
	return res.Error
}
