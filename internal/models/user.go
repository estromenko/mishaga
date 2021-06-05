package models

type User struct {
	ID         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	City       string `json:"city" db:"city"`
	Dorm       string `json:"dorm" db:"dorm"`
	RoomNumber *int   `json:"room_number" db:"room_number"`
}
