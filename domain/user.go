package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	FullName string
	Gender   string
	Birthday time.Time
	City     string
	Role     string
}

type IUserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetById(ctx context.Context, userId uuid.UUID) (*domain.User, error)
	GetAll(ctx context.Context, page int) ([]*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
