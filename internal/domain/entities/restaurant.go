package entities

import "time"

type Restaurant struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `validate:"required,min=3"`
	Address   string
	OwnerId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type UserLoginDto struct {
// 	Phone    string `validate:"required,min=3"`
// 	Password string `validate:"required,min=3"`
// }

// type UserUpdateDto struct {
// 	Id        uint   `validate:"omitempty"`
// 	Name      string `validate:"omitempty,min=3"`
// 	Surname   string `validate:"omitempty,min=3"`
// 	Phone     string `validate:"omitempty,min=3"`
// 	Email     string `validate:"omitempty,email"`
// 	Password  string `validate:"omitempty,min=3"`
// 	Token     string
// 	UpdatedAt time.Time
// }
