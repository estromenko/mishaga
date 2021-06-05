package service

import (
	"mishaga/internal/repo"
)

type Config struct {
	Salt string `json:"salt"`
}

type Services struct {
	UserService *UserService
}

func InitServices(repos *repo.Repositories, config *Config) *Services {
	return &Services{
		UserService: &UserService{repos: repos, config: config},
	}
}

