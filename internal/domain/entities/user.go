package entities

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3"`
	Surname   string    `json:"surname" validate:"required,min=3"`
	Phone     string    `json:"phone" validate:"required,min=3"`
	Email     string    `json:"email"`
	Password  string    `json:"password" validate:"required,min=3"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLoginDto struct {
	Phone    string `validate:"required,min=3"`
	Password string `validate:"required,min=3"`
}

type UserUpdateDto struct {
	Id        uint   `validate:"omitempty"`
	Name      string `validate:"omitempty,min=3"`
	Surname   string `validate:"omitempty,min=3"`
	Phone     string `validate:"omitempty,min=3"`
	Email     string `validate:"omitempty,email"`
	Password  string `validate:"omitempty,min=3"`
	Token     string
	UpdatedAt time.Time
}
