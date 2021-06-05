package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
	"github.com/rs/zerolog"

	"mishaga/pkg/database"

	jwtware "github.com/gofiber/jwt/v2"
)

type Config struct {
	Port      int    `json:"port"`
	JWTSecret string `json:"jwt_secret"`
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

	public := app.Group("/")
	{
		public.All("/", s.IndexHandler)
		public.All("/reg", s.RegistrationHandler)
		public.All("/login", s.LoginHandler)
	}

	private := app.Group("/")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(s.config.JWTSecret),
		ErrorHandler: s.NotAuthorizedHandler,
		TokenLookup:  "cookie:token",
	}))
	{
		private.Get("/main", s.MainHandler)
		private.Get("/ads", s.AdsHandler)
		private.Get("/theme", s.ThemeHandler)
		private.All("/new_theme", s.NewThemeHandler)
		private.All("/profile", s.ProfileHandler)
	}

	app.Use(s.NotFoundHandler)

	return app
}

func (s *Server) Run() error {
	app := s.route()
	return app.Listen(":" + os.Getenv("PORT"))
}
