package domain

import (
	"context"
	"github.com/dlankinl/bmstu-ppo-bl/domain"

	"github.com/google/uuid"
)

type UserSkill struct {
	UserId  uuid.UUID
	SkillId uuid.UUID
}

type IUserSkillRepository interface {
	Create(ctx context.Context, pair *domain.UserSkill) error
	Delete(ctx context.Context, pair *domain.UserSkill) error
	GetUserSkillsByUserId(ctx context.Context, userId uuid.UUID, page int) ([]*domain.UserSkill, error)
	GetUserSkillsBySkillId(ctx context.Context, skillId uuid.UUID, page int) ([]*domain.UserSkill, error)
}
