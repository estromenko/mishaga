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
	r.db.Get(&user, `SELECT * FROM users WHERE id = $1`, id)
	return user
}

func (r *UserRepo) GetByEmail(email string) *models.User {
	var user models.User
	r.db.Get(&user, `SELECT * FROM users WHERE email = $1`, email)
	return &user
}

func (r *UserRepo) SetUserImage(id int, avatar string) error {
	_, err := r.db.Exec(`UPDATE users SET avatar = $1 WHERE id = $2`, avatar, id)
	return err
}

func (r *UserRepo) UpdateUser(id int, user *models.User) error {
	_, err := r.db.Exec(
		`UPDATE users SET 
		first_name = $1, last_name = $2, city = $3,
		dorm = $4, room_number = $5, university = $6, email = $7 WHERE id = $8`,
		user.FirstName, user.LastName, user.City, user.Dorm,
		user.RoomNumber, user.University, user.Email, id,
	)
	return err
}
