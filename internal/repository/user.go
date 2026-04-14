package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `SELECT user_id, full_name, email, password, address, phone, profile_pic, created_at, role FROM users`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `SELECT user_id, full_name, email, password, address, phone, profile_pic, created_at, role FROM users WHERE email = $1`

	rows, err := r.DB.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Save(input models.UserRegisterInput) error {
	query := `INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(context.Background(), query, input.FullName, input.Email, input.Password)
	return err
}

func (r *UserRepository) UpdatePassword(email, newPassword string) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`

	_, err := r.DB.Exec(context.Background(), query, newPassword, email)
	return err
}

func (r *UserRepository) UpdateProfile(userID int, input models.UserUpdateInput) error {
	query := `
	UPDATE users
	SET full_name = $1,
	    email = $2,
	    phone = $3,
	    address = $4
	WHERE user_id = $5
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		input.FullName,
		input.Email,
		input.Phone,
		input.Address,
		userID,
	)

	return err
}

func (r *UserRepository) UpdateProfilePic(userID int, path string) error {
	query := `
	UPDATE users
	SET profile_pic = $1
	WHERE user_id = $2
	`

	_, err := r.DB.Exec(context.Background(), query, path, userID)
	return err
}
