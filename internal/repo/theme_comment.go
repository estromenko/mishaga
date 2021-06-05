package repo

import (
	"mishaga/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type ThemeCommentRepo struct {
	db *sqlx.DB
}

type FullThemeComment struct {
	Text      string `json:"text" db:"text" form:"text"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

func (t *ThemeCommentRepo) GetAllByThemeID(id int) []*models.ThemeComment {
	var comments []*models.ThemeComment
	t.db.Select(&comments, `SELECT * FROM theme_comments WHERE theme_id = $1`, id)
	return comments
}

func (t *ThemeCommentRepo) GetAllByThemeIDFull(id int) []*FullThemeComment {
	var comments []*FullThemeComment
	t.db.Select(&comments,
		`SELECT c.text, u.first_name, u.last_name 
		FROM theme_comments as c 
		JOIN users AS u ON c.owner_id = u.id 
		WHERE c.theme_id = $1 
		ORDER BY c.created_at`, id)
	return comments
}

func (t *ThemeCommentRepo) Create(comment *models.ThemeComment) error {
	return t.db.QueryRow(
		`INSERT INTO theme_comments (theme_id, owner_id, text, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		&comment.ThemeID, &comment.OwnerID, &comment.Text, time.Now(),
	).Scan(&comment.ID)
}
