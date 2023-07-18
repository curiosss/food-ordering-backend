package storage

import "template2/internal/domain/entities"

type Authorization interface {
	CreateUser(user *entities.User) (*entities.User, error)
	Update(user *entities.UserUpdateDto) (*entities.User, error)
	Delete(user_id uint) error
	GetUser(loginDto *entities.UserLoginDto) (*entities.User, error)
}

type Restaurants interface {
	Create()
}

type Departments interface {
	GetAll() ([]entities.Department, error)
}

type Storages struct {
	Authorization
	Restaurants
	Departments
}
