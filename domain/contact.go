package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

	"github.com/google/uuid"
)

type Contact struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
	Value   string
}

type IContactsRepository interface {
	Create(ctx context.Context, contact *domain.Contact) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Contact, error)
	GetByOwnerId(ctx context.Context, id uuid.UUID, page int) ([]*domain.Contact, error)
	Update(ctx context.Context, contact *domain.Contact) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
