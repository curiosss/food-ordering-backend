package entities

import "time"

type Department struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3"`
	ImageUrl  string    `json:"imageUrl"`
	IsEnabled bool      `json:"isEnabled"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
