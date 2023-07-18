package main

import (
	"fmt"
	handler "template2/internal/app/handlers"
	"template2/internal/app/services"
	"template2/internal/app/storages"
	"template2/internal/middleware"

	// logmiddleware "template2/internal/middleware"
	"template2/pkg/logger"
	"template2/pkg/postgresql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {

	logger, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		panic("logger didn't initialized")
	}
	logger.Info("Logger initialized")

	db, err := postgresql.NewPostgresDB()
	if err != nil {
		logger.Fatalf("error initing db: %s", err.Error())
	}

	err = postgresql.SyncDb(db)
	if err != nil {
		logger.Fatalf("error syncing db: %s", err.Error())
	}

	postgresql.SeedConstants(db)

	storages := storages.NewStorage(db)
	services := services.NewServices(storages)
	handlers := handler.NewHandler(services, logger, storages)

	app := fiber.New()

	logMiddleware := middleware.NewLogMiddleware(logger)

	app.Use(adaptor.HTTPMiddleware(logMiddleware.Log))
	app.Use(middleware.AuthMIddleware)
	// app.Use(logger)

	app.Static("/public", "./public")

	handlers.InitRoutes(app)
	app.Get("/public/metrics", monitor.New(monitor.Config{Title: "Template 2"}))

	app.Listen(":3000")
	fmt.Println(storages)
	fmt.Println(services)

}
