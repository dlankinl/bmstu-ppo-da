package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

	"github.com/google/uuid"
)

type ActivityField struct {
	ID          uuid.UUID
	Name        string
	Description string
	Cost        float32
}

type IActivityFieldRepository interface {
	Create(context.Context, *domain.ActivityField) error
	DeleteById(context.Context, uuid.UUID) error
	Update(context.Context, *domain.ActivityField) error
	GetById(context.Context, uuid.UUID) (*domain.ActivityField, error)
	GetMaxCost(context.Context) (float32, error)
	GetAll(context.Context, int) ([]*domain.ActivityField, error)
}
