package repo

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepo *UserRepo
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepo: &UserRepo{db: db},
	}
}
