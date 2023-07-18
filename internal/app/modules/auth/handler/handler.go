package auth

import (
	"fmt"
	"strconv"
	"template2/internal/domain/entities"
	"template2/internal/domain/services"
	"template2/internal/domain/storage"
	"template2/internal/middleware"
	validator "template2/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Authorization interface {
	SignUp(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	UpdateAccount(ctx *fiber.Ctx) error
	DeleteAccount(ctx *fiber.Ctx) error
}

type AuthorizationImpl struct {
	serv   *services.Services
	logger *logrus.Logger
	repo   storage.Authorization
}

func NewAuthorization(service *services.Services, logger *logrus.Logger, storage storage.Authorization) Authorization {
	return &AuthorizationImpl{serv: service, logger: logger, repo: storage}
}

func (h *AuthorizationImpl) SignUp(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := validator.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	// err := h.serv.Authorization.SignUp(user)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"status":  false,
	// 		"message": err.Error(),
	// 	})
	// }
	dbUser, err := h.repo.GetUser(&entities.UserLoginDto{Phone: user.Phone})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
		})
	}
	if dbUser != nil && dbUser.Id > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "User with this phone number already exists",
		})
	}

	user, err = h.repo.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
		})
	}
	tokenStr := middleware.GenerateToken(user)
	user.Token = tokenStr
	user, err = h.repo.Update(&entities.UserUpdateDto{Id: user.Id, Token: user.Token})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

func (h *AuthorizationImpl) Login(c *fiber.Ctx) error {
	// fmt.Println("login handler started:")
	loginDto := new(entities.UserLoginDto)
	if err := c.BodyParser(loginDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := validator.ValidateStruct(loginDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	user, err := h.serv.Authorization.Login(loginDto)
	// fmt.Println("login handler after service:")
	// fmt.Println("login handler after service error:", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   user,
	})

}

func (h *AuthorizationImpl) Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

func (h *AuthorizationImpl) UpdateAccount(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Get("userId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	userUpdateDto := new(entities.UserUpdateDto)
	userUpdateDto.Id = uint(userId)
	if err := c.BodyParser(userUpdateDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := validator.ValidateStruct(userUpdateDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	user, err := h.serv.Authorization.Update(userUpdateDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

func (h *AuthorizationImpl) DeleteAccount(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Get("userId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	err = h.serv.Authorization.Delete(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"errors": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": fmt.Sprintf("user deleted successfully with id: %d", userId)})

}
