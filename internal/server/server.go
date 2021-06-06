package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
	"github.com/rs/zerolog"

	"mishaga/internal/repo"
	"mishaga/internal/service"
	"mishaga/pkg/database"

	jwtware "github.com/gofiber/jwt/v2"
)

type Config struct {
	Port      int             `json:"port"`
	JWTSecret string          `json:"jwt_secret"`
	Services  *service.Config `json:"services"`
}

type Server struct {
	config   *Config
	db       *database.Database
	logger   *zerolog.Logger
	repos    *repo.Repositories
	services *service.Services
}

func NewServer(config *Config, db *database.Database, logger *zerolog.Logger) *Server {
	repos := repo.InitRepositories(db.DB())
	services := service.InitServices(repos, config.Services)
	return &Server{
		config:   config,
		db:       db,
		logger:   logger,
		repos:    repos,
		services: services,
	}
}

func (s *Server) route() *fiber.App {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "static")
	app.Static("/uploads", "uploads")

	app.Use(favicon.New())
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
		private.All("/themes", s.ThemeHandler)
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
