package repo

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepo         *UserRepo
	ThemeRepo        *ThemeRepo
	ThemeCommentRepo *ThemeCommentRepo
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepo:         &UserRepo{db: db},
		ThemeRepo:        &ThemeRepo{db: db},
		ThemeCommentRepo: &ThemeCommentRepo{db: db},
	}
}
