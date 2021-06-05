package models

type Theme struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title" form:"title"`
	Category    string `json:"category" db:"category" form:"category"`
	Description string `json:"description" db:"description" form:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
