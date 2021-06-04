package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Redirect("reg")
}

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
	if c.Method() == "GET" {
		return c.Render("new_theme", fiber.Map{})
	}
	return c.Redirect("/main")
}

func ProfileHandler(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{})
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{})
}
