package server

import (
	"fmt"
	"mishaga/internal/models"
	"strconv"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) IndexHandler(c *fiber.Ctx) error {
	return c.Redirect("reg")
}

func (s *Server) RegistrationHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		c.ClearCookie("token")
		return c.Render("reg", fiber.Map{})
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.SendString("Error parsing body: " + err.Error())
	}

	if err := s.services.UserService.Create(&user); err != nil {
		return c.SendString(err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	t, _ := token.SignedString([]byte(s.config.JWTSecret))

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	return c.Redirect("/main")
}

func (s *Server) LoginHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		c.ClearCookie("token")
		return c.Render("login", fiber.Map{})
	}

	user := s.repos.UserRepo.GetByEmail(c.FormValue("email"))
	if user == nil {
		return c.Redirect("/login")
	}

	if !s.services.UserService.ComparePasswords(user, c.FormValue("password")) {
		return c.Render("login", fiber.Map{})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	t, _ := token.SignedString([]byte(s.config.JWTSecret))

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	return c.Redirect("/main")
}

func (s *Server) MainHandler(c *fiber.Ctx) error {
	themes := s.repos.ThemeRepo.GetAll()
	return c.Render("main", fiber.Map{
		"themes": themes,
	})
}

func (s *Server) AdsHandler(c *fiber.Ctx) error {
	return c.Render("ads", fiber.Map{})
}

func (s *Server) ThemeHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))
	if c.Method() == "GET" {
		comments := s.repos.ThemeCommentRepo.GetAllByThemeIDFull(id)
		return c.Render("theme", fiber.Map{
			"comments": comments,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	user := s.repos.UserRepo.GetByEmail(email)
	if user == nil {
		s.logger.Debug().Msg(email)
		return c.Redirect("/login")
	}

	var comment models.ThemeComment
	c.BodyParser(&comment)
	comment.OwnerID = user.ID
	comment.ThemeID = id
	s.repos.ThemeCommentRepo.Create(&comment)

	path := fmt.Sprintf("/themes?id=%d", id)
	return c.Redirect(path)
}

func (s *Server) NewThemeHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("new_theme", fiber.Map{})
	}

	var theme models.Theme
	c.BodyParser(&theme)
	s.repos.ThemeRepo.Create(&theme)

	return c.Redirect("/main")
}

func (s *Server) ProfileHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		user := s.repos.UserRepo.GetByEmail(email)
		if user == nil {
			s.logger.Debug().Msg(email)
			return c.Redirect("/login")
		}

		return c.Render("profile", fiber.Map{
			"user": user,
		})
	}

	return c.Render("profile", fiber.Map{})
}

func (s *Server) NotFoundHandler(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{})
}

func (s *Server) NotAuthorizedHandler(c *fiber.Ctx, err error) error {
	return c.Redirect("/login")
}
