package postgtres

import (
	"context"
	"eldeck/eldeck/internal/auth/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) CreateUser(ctx context.Context, user *models.RegisterParams, hashedPassword string) error {
	const query = `
INSERT INTO users.user (name, login, password_hash)
VALUES ($1, $2, $3);
`
	_, err := repo.db.ExecContext(ctx, query, user.Name, user.Login, hashedPassword)
	if err != nil {
		slog.Error("failed execute query for create user", "error", err)
		return fmt.Errorf("failed exeute query for create user: %w", err)
	}

	return nil
}

func (repo *AuthRepository) GetUserPassword(ctx context.Context, login string) (string, error) {
	const query = `
SELECT password_hash FROM users.user WHERE login = $1 AND blocked = FALSE;
`
	var password string
	err := repo.db.GetContext(ctx, &password, query, login)
	if err != nil {
		slog.Error("failed execute query for validate user", "error", err)
		return "", fmt.Errorf("failed exeute query for validate user: %w", err)
	}

	return password, nil
}
