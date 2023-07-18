package departments

import (
	"template2/internal/domain/entities"
	"template2/internal/domain/storage"

	"gorm.io/gorm"
)

type DepartmentStorageImpl struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) storage.Departments {
	return &DepartmentStorageImpl{db: db}
}

func (s *DepartmentStorageImpl) GetAll() ([]entities.Department, error) {
	users := []entities.Department{}
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
