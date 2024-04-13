package postgtres

import (
	"context"
	"eldick/eldick/internal/auth/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.RegisterParams, hashedPassword string) error
	GetUserPassword(ctx context.Context, login string) (string, error)
}
