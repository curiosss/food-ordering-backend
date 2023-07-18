package departments

import (
	"template2/internal/domain/storage"
	"template2/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DepartmentHandler interface {
	GetAll(ctx *fiber.Ctx) error
}

type DepartmentHandlerImpl struct {
	logger *logrus.Logger
	repo   storage.Departments
}

func NewDepartmentsHandler(logger *logrus.Logger, repo storage.Storages) DepartmentHandler {
	return &DepartmentHandlerImpl{logger: logger, repo: repo}
}

func (h *DepartmentHandlerImpl) GetAll(c *fiber.Ctx) error {

	departments, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"error":  err,
		})
	}

	// dd := map[string]any{
	// 	"status": true,
	// 	"data": map[string]any{
	// 		"departments": departments,
	// 	}}
	res := utils.Response{Status: true, Data: map[string]any{"departments": departments}}

	return c.Status(fiber.StatusOK).JSON(res)
}
