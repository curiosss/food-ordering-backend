package entities

import "time"

type Roles struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UserRoles struct {
	UserId         uint      `json:"userId"`
	RoleId         uint      `json:"roleId"`
	Status         string    `json:"status"`
	ExpirationDate time.Time `json:"expirationDate"`
}
