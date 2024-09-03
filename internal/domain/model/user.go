package model

import (
	"errors"
	"strings"

	"github.com/ruslanguns/go-chat/internal/domain"
)

type User struct {
	domain.BaseEntity
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
}

func NewUser(username, email, password string) (*User, error) {
	u := &User{
		BaseEntity: domain.BaseEntity{},
		Username:   strings.TrimSpace(username),
		Email:      strings.TrimSpace(email),
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}

func (u *User) ChangeEmail(newEmail string) error {
	newEmail = strings.TrimSpace(newEmail)
	if newEmail == "" {
		return errors.New("email cannot be empty")
	}
	u.Email = newEmail
	return nil
}

func (u *User) ChangeUsername(newUsername string) error {
	newUsername = strings.TrimSpace(newUsername)
	if newUsername == "" {
		return errors.New("username cannot be empty")
	}
	u.Username = newUsername
	return nil
}
