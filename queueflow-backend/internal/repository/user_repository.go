package repository

import (
	"database/sql"
	"errors"

	"queueflow/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create user
func (r *UserRepository) Create(
	user models.User,
) error {

	query := `
	INSERT INTO users
	(
		id,
		username,
		email,
		password,
		role,
		created_at
	)
	VALUES
	($1,$2,$3,$4,$5,NOW())
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

// Find by email
func (r *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {

	query := `
	SELECT
		id,
		username,
		email,
		password,
		role,
		created_at,
		last_login
	FROM users
	WHERE email=$1
	`

	user := &models.User{}

	err := r.db.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.LastLogin,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}

// Find by ID
func (r *UserRepository) FindByID(
	id string,
) (*models.User, error) {

	query := `
	SELECT
		id,
		username,
		email,
		role,
		created_at,
		last_login
	FROM users
	WHERE id=$1
	`

	user := &models.User{}

	err := r.db.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
		&user.LastLogin,
	)

	return user, err
}

// Update last login
func (r *UserRepository) UpdateLastLogin(
	id string,
) error {

	_, err := r.db.Exec(
		`
		UPDATE users
		SET last_login=NOW()
		WHERE id=$1
		`,
		id,
	)

	return err
}
