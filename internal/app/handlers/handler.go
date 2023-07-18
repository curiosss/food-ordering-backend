package handler

import (
	"fmt"
	auth "template2/internal/app/modules/auth/handler"
	"template2/internal/app/modules/departments"
	"template2/internal/domain/services"
	"template2/internal/domain/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	auth.Authorization
	departments.DepartmentHandler
}

func NewHandler(sv *services.Services, logger *logrus.Logger, storage *storage.Storages) *Handler {
	return &Handler{Authorization: auth.NewAuthorization(sv, logger, storage.Authorization), DepartmentHandler: departments.NewDepartmentsHandler(logger, *storage)}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	fmt.Println("initing routes")

	public := app.Group("/public")

	publicUser := app.Group("/public/user")
	publicUser.Post("/login", h.Authorization.Login)
	publicUser.Post("/sign-up", h.Authorization.SignUp)

	user := app.Group("/user")
	user.Get("/logout", h.Authorization.Logout)
	user.Put("/update", h.Authorization.UpdateAccount)
	user.Get("/delete", h.Authorization.DeleteAccount)

	public.Get("/departments", h.DepartmentHandler.GetAll)

}
