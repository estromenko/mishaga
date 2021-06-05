package service

import (
	"encoding/base64"
	"fmt"
	"mishaga/internal/models"
	"mishaga/internal/repo"

	"golang.org/x/crypto/argon2"
)

type UserService struct {
	repos  *repo.Repositories
	config *Config
}

func (u *UserService) hashPassword(password string) string {
	// Helper struct for password hashing
	type passwordConfig struct {
		time    uint32
		memory  uint32
		threads uint8
		keyLen  uint32
	}

	c := &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	hash := argon2.IDKey(
		[]byte(password),
		[]byte(u.config.Salt),
		c.time,
		c.memory,
		c.threads,
		c.keyLen,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString([]byte(u.config.Salt))
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
}

func (u *UserService) ComparePasswords(user *models.User, pass string) bool {
	return u.hashPassword(pass) == user.Password
}

func (u *UserService) Create(user *models.User) error {
	user.Password = u.hashPassword(user.Password)
	return u.repos.UserRepo.Create(user)
}
