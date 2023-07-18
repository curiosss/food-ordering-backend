package storages

import (
	auth "template2/internal/app/modules/auth/storage"
	"template2/internal/app/modules/departments"
	"template2/internal/domain/storage"

	"gorm.io/gorm"
)

func NewStorage(db *gorm.DB) *storage.Storages {
	return &storage.Storages{
		Authorization: auth.NewAuthorizationImp(db),
		Departments:   departments.NewStorage(db),
	}
}
