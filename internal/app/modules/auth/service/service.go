package auth

import (
	"errors"
	"template2/internal/domain/entities"
	service "template2/internal/domain/services"
	"template2/internal/domain/storage"
)

type AuthorizationImpl struct {
	repo storage.Authorization
}

func NewAuthService(repo storage.Authorization) service.Authorization {
	return &AuthorizationImpl{repo: repo}
}

func (r *AuthorizationImpl) SignUp(user *entities.User) (*entities.User, error) {
	user, err := r.repo.GetUser(&entities.UserLoginDto{Phone: user.Phone})
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("You already signed up with this phone number")
	}
	return r.repo.CreateUser(user)
}

func (r *AuthorizationImpl) Login(loginDto *entities.UserLoginDto) (*entities.User, error) {
	user, err := r.repo.GetUser(loginDto)
	// fmt.Println("login service:")
	// fmt.Println(user)
	// tokenstr := middleware.GenerateToken(user)
	// fmt.Println(tokenstr)
	if err != nil {
		return nil, err
	}
	if user.Password != loginDto.Password {
		return nil, errors.New("user credentials didn't match")
	}
	return user, nil
}

func (r *AuthorizationImpl) Logout(user_id uint) error {
	return nil
}

func (r *AuthorizationImpl) Update(user *entities.UserUpdateDto) (*entities.User, error) {
	return r.repo.Update(user)
}

func (r *AuthorizationImpl) Delete(user_id uint) error {
	return r.repo.Delete(user_id)
}
