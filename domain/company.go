package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

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
	Create(ctx context.Context, company *domain.Company) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Company, error)
	GetByOwnerId(ctx context.Context, id uuid.UUID, page int) ([]*domain.Company, error)
	GetAll(ctx context.Context, page int) ([]*domain.Company, error)
	Update(ctx context.Context, company *domain.Company) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
