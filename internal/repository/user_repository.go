package repository

import (
	"strings"

	"github.com/ruslanguns/go-chat/internal/domain"
	"github.com/ruslanguns/go-chat/internal/domain/model"
	"github.com/ruslanguns/go-chat/internal/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id domain.EntityID) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id domain.EntityID) error
	List(offset, limit int) ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.username") {
			return errors.NewAppError(errors.ErrAlreadyExists, "A user with this username already exists")
		}
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.email") {
			return errors.NewAppError(errors.ErrAlreadyExists, "A user with this email already exists")
		}
		return errors.NewAppError(errors.ErrInternal, err.Error())
	}
	return nil
}

func (r *userRepository) GetByID(id domain.EntityID) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id.String()).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError(errors.ErrNotFound, "User not found")
		}
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get user")
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError(errors.ErrNotFound, "User not found")
		}
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get user")
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError(errors.ErrNotFound, "User not found")
		}
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get user")
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to update user")
	}
	return nil
}

func (r *userRepository) Delete(id domain.EntityID) error {
	result := r.db.Delete(&model.User{}, "id = ?", id.String())
	if result.Error != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to delete user")
	}
	if result.RowsAffected == 0 {
		return errors.NewAppError(errors.ErrNotFound, "User not found")
	}
	return nil
}

func (r *userRepository) List(offset, limit int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to list users")
	}
	return users, nil
}
