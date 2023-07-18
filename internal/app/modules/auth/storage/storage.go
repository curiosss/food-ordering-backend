package auth

import (
	"fmt"
	"template2/internal/domain/entities"
	"template2/internal/domain/storage"
	"gorm.io/gorm"
)

type AuthorizationImpl struct {
	db *gorm.DB
}

func NewAuthorizationImp(db *gorm.DB) storage.Authorization {
	return &AuthorizationImpl{db: db}
}

func (r *AuthorizationImpl) CreateUser(user *entities.User) (*entities.User, error) {

	fmt.Println("signin in repository : ", user)
	result := r.db.Create(user).Scan(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *AuthorizationImpl) Update(userUpdateDto *entities.UserUpdateDto) (*entities.User, error) {
	// result := r.db.Save(user)
	fmt.Println("inside storage users update")
	fmt.Println(userUpdateDto)
	user := new(entities.User)
	result := r.db.Table("users").Updates(userUpdateDto).Scan(user)
	sql := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("users").Updates(userUpdateDto)
	})
	fmt.Println(sql)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}

func (r *AuthorizationImpl) Delete(user_id uint) error {
	sql := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("users").Where("id", user_id).Delete(&user_id)
	})
	fmt.Println(sql)

	result := r.db.Table("users").Where("id", user_id).Delete(&user_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthorizationImpl) GetUser(loginDto *entities.UserLoginDto) (*entities.User, error) {
	// res := new([]entities.User)
	res := new(entities.User)

	// sql := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Table("users").Where("phone=?", "99362216478").Find(res)
	// })
	// fmt.Println(sql)

	// result := r.db.Model(&entities.User{Phone: "99362216478"}).Scan(res)
	result := r.db.Table("users").Where("phone=?", loginDto.Phone).Find(res)
	if result.Error != nil {
		return nil, result.Error
	}
	// fmt.Println(result.RowsAffected)
	// fmt.Println(res)
	return res, nil
}
