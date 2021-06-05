package models

import "time"

type ThemeComment struct {
	ID        int       `json:"id" db:"id"`
	ThemeID   int       `json:"theme_id" db:"theme_id"`
	OwnerID   int       `json:"owner_id" db:"owner_id"`
	Text      string    `json:"text" db:"text" form:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
