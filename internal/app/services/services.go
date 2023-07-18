package services

import (
	auth "template2/internal/app/modules/auth/service"
	"template2/internal/domain/services"
	"template2/internal/domain/storage"
)

func NewServices(storages *storage.Storages) *services.Services {
	return &services.Services{
		Authorization: auth.NewAuthService(storages.Authorization),
	}
}
