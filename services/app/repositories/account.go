package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Account, error)
	FindByProvider(ctx context.Context, providerType, providerID, providerAccountID string) (*Account, error)
	Insert(ctx context.Context, account *Account) error
}

type Account struct {
	ID                uuid.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ProviderType      string
	ProviderID        string
	ProviderAccountID string
	UserID            uuid.UUID
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (a *accountRepository) FindByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	const query = `
		SELECT id, created_at, updated_at, provider_type, provider_id, provider_account_id, user_id FROM accounts
		WHERE id = $1
	`
	row := a.db.QueryRowContext(ctx, query, id)
	account := Account{}
	if err := row.Scan(
		&account.ID,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.ProviderType,
		&account.ProviderID,
		&account.ProviderAccountID,
		&account.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

func (a *accountRepository) FindByProvider(ctx context.Context, providerType, providerID, providerAccountID string) (*Account, error) {
	const query = `
		SELECT id, created_at, updated_at, provider_type, provider_id, provider_account_id, user_id FROM accounts
		WHERE provider_type = $1
		AND provider_id = $2
		AND provider_account_id = $3
	`
	row := a.db.QueryRowContext(ctx, query, providerType, providerID, providerAccountID)
	account := Account{}
	if err := row.Scan(
		&account.ID,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.ProviderType,
		&account.ProviderID,
		&account.ProviderAccountID,
		&account.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

func (a *accountRepository) Insert(ctx context.Context, account *Account) error {
	const query = `
		INSERT INTO accounts (id, created_at, updated_at, provider_type, provider_id, provider_account_id, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, err := a.db.ExecContext(
		ctx,
		query,
		account.ID,
		account.CreatedAt,
		account.UpdatedAt,
		account.ProviderType,
		account.ProviderID,
		account.ProviderAccountID,
		account.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}
