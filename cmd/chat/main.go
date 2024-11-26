package main

import (
	"time"

	"github.com/karalarmehmet/chatapplication/internal/app"
	"github.com/karalarmehmet/chatapplication/internal/db"
	"github.com/karalarmehmet/chatapplication/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("web/views", ".tmpl")
	ap := fiber.New(fiber.Config{
		AppName:           "Chat",
		Views:             engine,
		ViewsLayout:       "layouts/main",
		ErrorHandler:      handlers.CustomErrorHandler,
		PassLocalsToViews: true,
	})
	db, storage := db.InitDb()
	ss := session.New(session.Config{
		Storage:    storage,
		Expiration: time.Minute * 60,
	})

	// Dependency injection
	newApp := app.NewApp(ap, db, ss)

	newApp.App.Use(logger.New())

	handlers.SetupRoutes(newApp)

	newApp.Start()
}
