package services

import (
	"eliftech_1/models"
	"eliftech_1/repositories"
)

type PostService interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id uint) (*models.Post, error)
	CreatePost(post *models.Post) (*models.Post, error)
	DeletePost(id uint) error
}

type postService struct {
	postRepository repositories.PostRepository
}

func (s *postService) GetPosts() ([]models.Post, error) {
	res, err := s.postRepository.GetPosts()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *postService) GetPostById(id uint) (*models.Post, error) {
	res, err := s.postRepository.GetPostById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	res, err := s.postRepository.CreatePost(post)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *postService) DeletePost(id uint) error {
	err := s.postRepository.DeletePost(id)

	if err != nil {
		return err
	}

	return nil
}
