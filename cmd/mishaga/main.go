package main

import (
	"mishaga/internal/server"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "static")
	app.All("/reg", server.RegistrationHandler)
	app.All("/login", server.LoginHandler)
	app.Get("/main", server.MainHandler)
	app.Get("/ads", server.AdsHandler)
	app.Get("/theme", server.ThemeHandler)
	app.Get("/new_theme", server.NewThemeHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
