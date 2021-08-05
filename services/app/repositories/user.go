package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByPrimaryEmail(ctx context.Context, primaryEmail string) (*User, error)
	Insert(ctx context.Context, user *User) error
}

type User struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PrimaryEmail string
	FullName     string
	AvatarURL    string
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	const query = `
		SELECT id, created_at, updated_at, primary_email, full_name, avatar_url FROM users
		WHERE id = $1
	`
	row := u.db.QueryRowContext(ctx, query, id)
	user := User{}
	if err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PrimaryEmail,
		&user.FullName,
		&user.AvatarURL,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByPrimaryEmail(ctx context.Context, primaryEmail string) (*User, error) {
	const query = `
		SELECT id, created_at, updated_at, primary_email, full_name, avatar_url FROM users
		WHERE primary_email = $1
	`
	row := u.db.QueryRowContext(ctx, query, primaryEmail)
	user := User{}
	if err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PrimaryEmail,
		&user.FullName,
		&user.AvatarURL,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Insert(ctx context.Context, user *User) error {
	const query = `
		INSERT INTO users (id, created_at, updated_at, primary_email, full_name, avatar_url)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := u.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.PrimaryEmail,
		user.FullName,
		user.AvatarURL,
	)
	if err != nil {
		return err
	}
	return nil
}
