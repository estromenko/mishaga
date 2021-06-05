package repo

import (
	"mishaga/internal/models"

	"github.com/jmoiron/sqlx"
)

type ThemeRepo struct {
	db *sqlx.DB
}

func (t *ThemeRepo) GetByID(id int) *models.Theme {
	var theme *models.Theme
	t.db.QueryRow(`SELECT * FROM themes WHERE id = $1`, id).Scan(
		&theme.ID, &theme.Title, &theme.Category, &theme.Description)
	return theme
}

func (t *ThemeRepo) GetAll() []*models.Theme {
	var themes []*models.Theme
	t.db.Select(&themes, `SELECT * FROM themes`)
	return themes
}

func (t *ThemeRepo) Create(theme *models.Theme) error {
	return t.db.QueryRow(
		`INSERT INTO themes (title, category, description)
		VALUES ($1, $2, $3) RETURNING id`,
		&theme.Title, &theme.Category, &theme.Description,
	).Scan(&theme.ID)
}
