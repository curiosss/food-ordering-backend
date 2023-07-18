package services

import (
	"template2/internal/domain/entities"
)

type Authorization interface {
	SignUp(user *entities.User) (*entities.User, error)
	Login(user *entities.UserLoginDto) (*entities.User, error)
	Logout(user_id uint) error
	Update(user *entities.UserUpdateDto) (*entities.User, error)
	Delete(user_id uint) error
}

type Restaurants interface {
}

type Services struct {
	Authorization
}
