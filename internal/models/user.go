package models

type User struct {
	ID         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email" form:"email"`
	FirstName  string `json:"first_name" db:"first_name" form:"first_name"`
	LastName   string `json:"last_name" db:"last_name" form:"last_name"`
	City       string `json:"city" db:"city" form:"city"`
	Dorm       string `json:"dorm" db:"dorm" form:"dorm"`
	RoomNumber *int   `json:"room_number" db:"room_number" form:"room_number"`
	University string `json:"university" db:"university" form:"university"`
	Password   string `json:"password" db:"password" form:"password"`
}
