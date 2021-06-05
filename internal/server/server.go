package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
	"github.com/rs/zerolog"

	"mishaga/pkg/database"
)

type Config struct {
	Port int `json:"port"`
}

type Server struct {
	config *Config
	db     *database.Database
	logger *zerolog.Logger
}

func NewServer(config *Config, db *database.Database, logger *zerolog.Logger) *Server {
	return &Server{
		config: config,
		db:     db,
		logger: logger,
	}
}

func (s *Server) route() *fiber.App {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "static")

	app.Use(logger.New())

	app.All("/", s.IndexHandler)
	app.All("/reg", s.RegistrationHandler)
	app.All("/login", s.LoginHandler)
	app.Get("/main", s.MainHandler)
	app.Get("/ads", s.AdsHandler)
	app.Get("/theme", s.ThemeHandler)
	app.All("/new_theme", s.NewThemeHandler)
	app.All("/profile", s.ProfileHandler)

	app.Use(s.NotFoundHandler)

	return app
}

func (s *Server) Run() error {
	app := s.route()
	return app.Listen(":" + os.Getenv("PORT"))
}
