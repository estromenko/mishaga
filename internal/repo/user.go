package repo

import (
	"mishaga/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func (r *UserRepo) Create(user *models.User) error {
	return r.db.QueryRow(
		`INSERT INTO users 
		(email, first_name, last_name, city, dorm, room_number, university, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`,
		user.Email, user.FirstName, user.LastName, user.City, user.Dorm, user.RoomNumber, user.University, user.Password,
	).Scan(&user.ID)
}

func (r *UserRepo) GetByID(id int) *models.User {
	var user *models.User
	r.db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.City, &user.Dorm, &user.RoomNumber, &user.University, &user.Password)
	return user
}

func (r *UserRepo) GetByEmail(email string) *models.User {
	var user models.User
	r.db.QueryRow(`SELECT * FROM users WHERE email = $1`, email).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.City, &user.Dorm, &user.RoomNumber, &user.University, &user.Password)
	return &user
}
