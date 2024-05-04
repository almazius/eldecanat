package postgtres

import (
	"context"
	"eldeck/eldeck/internal/auth/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.RegisterParams, hashedPassword string) error
	GetUserPassword(ctx context.Context, login string) (string, error)
}
