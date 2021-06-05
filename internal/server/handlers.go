package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) IndexHandler(c *fiber.Ctx) error {
	return c.Redirect("reg")
}

func (s *Server) RegistrationHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("reg", fiber.Map{})
	}

	fmt.Println("User with email \"" + c.FormValue("email") + "\" entered")

	return c.Redirect("/main")
}

func (s *Server) LoginHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("login", fiber.Map{})
	}

	fmt.Println("User with email \"" + c.FormValue("email") + "\" entered")
	return c.Redirect("/main")
}

func (s *Server) MainHandler(c *fiber.Ctx) error {
	return c.Render("main", fiber.Map{})
}

func (s *Server) AdsHandler(c *fiber.Ctx) error {
	return c.Render("ads", fiber.Map{})
}

func (s *Server) ThemeHandler(c *fiber.Ctx) error {
	return c.Render("theme", fiber.Map{})
}

func (s *Server) NewThemeHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("new_theme", fiber.Map{})
	}
	return c.Redirect("/main")
}

func (s *Server) ProfileHandler(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{})
}

func (s *Server) NotFoundHandler(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{})
}
