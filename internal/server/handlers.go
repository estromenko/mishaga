package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegistrationHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("reg", fiber.Map{})
	}
	fmt.Println("User with email \"" + c.FormValue("email") + "\" entered")
	return c.Redirect("/main")
}

func LoginHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("login", fiber.Map{})
	}

	fmt.Println("User with email \"" + c.FormValue("email") + "\" entered")
	return c.Redirect("/main")
}

func MainHandler(c *fiber.Ctx) error {
	return c.Render("main", fiber.Map{})
}

func AdsHandler(c *fiber.Ctx) error {
	return c.Render("ads", fiber.Map{})
}

func ThemeHandler(c *fiber.Ctx) error {
	return c.Render("theme", fiber.Map{})
}

func NewThemeHandler(c *fiber.Ctx) error {
	return c.Render("new_theme", fiber.Map{})
}
