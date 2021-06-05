package repo

import (
	"mishaga/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ThemeRepo struct {
	db *sqlx.DB
}

type ThemeWithCommentQuantity struct {
	models.Theme
	Quantity int `json:"quantity" db:"quantity"`
}

func (t *ThemeRepo) GetByID(id int) *models.Theme {
	var theme *models.Theme
	t.db.QueryRow(`SELECT * FROM themes WHERE id = $1`, id).Scan(
		&theme.ID, &theme.Title, &theme.Category, &theme.Description, &theme.CreatedAt)
	return theme
}

func (t *ThemeRepo) GetAll() []*models.Theme {
	var themes []*models.Theme
	t.db.Select(&themes, `SELECT * FROM themes ORDER BY created_at`)
	return themes
}

func (t *ThemeRepo) GetAllWithCommentQuantity() []*ThemeWithCommentQuantity {
	var themes []*ThemeWithCommentQuantity
	t.db.Select(&themes,
		`SELECT t.*, COUNT(c.theme_id) AS quantity FROM themes AS t LEFT JOIN theme_comments AS c ON t.id = c.theme_id 
		GROUP BY t.id ORDER BY t.created_at DESC`)
	return themes
}

func (t *ThemeRepo) Create(theme *models.Theme) error {
	return t.db.QueryRow(
		`INSERT INTO themes (title, category, description, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		&theme.Title, &theme.Category, &theme.Description, time.Now(),
	).Scan(&theme.ID)
}
