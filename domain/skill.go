package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

	"github.com/google/uuid"
)

type Skill struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type ISkillRepository interface {
	Create(ctx context.Context, skill *domain.Skill) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Skill, error)
	GetAll(ctx context.Context, page int) ([]*domain.Skill, error)
	Update(ctx context.Context, skill *domain.Skill) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
