package services

import (
	"eliftech_1/models"
	"eliftech_1/repositories"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdteUser(id uint, user *models.CreateUserInput) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func (s *userService) GetUsers() ([]models.User, error) {
	res, err := s.userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) GetUserById(id uint) (*models.User, error) {
	res, err := s.userRepository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	res, err := s.userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) UpdteUser(id uint, user *models.CreateUserInput) error {
	err := s.userRepository.UpdteUser(id, user)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id uint) error {
	err := s.userRepository.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
