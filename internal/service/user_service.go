package service

import (
	"github.com/ruslanguns/go-chat/internal/domain"
	"github.com/ruslanguns/go-chat/internal/domain/model"
	"github.com/ruslanguns/go-chat/internal/errors"
	"github.com/ruslanguns/go-chat/internal/repository"
)

type UserService interface {
	CreateUser(username, email string) (*model.User, error)
	GetUserByID(id domain.EntityID) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id domain.EntityID) error
	ListUsers(offset, limit int) ([]*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(username, email string) (*model.User, error) {
	user, err := model.NewUser(username, email, "")
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Invalid user data")
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id domain.EntityID) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetByUsername(username)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) UpdateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return errors.NewAppError(errors.ErrInvalidInput, "Invalid user data")
	}
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id domain.EntityID) error {
	return s.userRepo.Delete(id)
}

func (s *userService) ListUsers(offset, limit int) ([]*model.User, error) {
	return s.userRepo.List(offset, limit)
}
