package domain

import (
	"context"

	"github.com/google/uuid"
)

type Company struct {
	ID              uuid.UUID
	OwnerID         uuid.UUID
	ActivityFieldId uuid.UUID
	Name            string
	City            string
}

type ICompanyRepository interface {
	Create(ctx context.Context, company *Company) error
	GetById(ctx context.Context, id uuid.UUID) (*Company, error)
	GetByOwnerId(ctx context.Context, id uuid.UUID, page int) ([]*Company, error)
	GetAll(ctx context.Context, page int) ([]*Company, error)
	Update(ctx context.Context, company *Company) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}