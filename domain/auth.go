package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

	"github.com/google/uuid"
)

type UserAuth struct {
	ID         uuid.UUID
	Username   string
	Password   string
	HashedPass string
	Role       string
}

type IAuthRepository interface {
	Register(ctx context.Context, authInfo *domain.UserAuth) (err error)
	GetByUsername(ctx context.Context, username string) (*domain.UserAuth, error)
}
