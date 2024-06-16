package domain

import (
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/google/uuid"
)

type IInteractor interface {
	GetMostProfitableCompany(period *domain.Period, companies []*domain.Company) (*domain.Company, error)
	CalculateUserRating(id uuid.UUID) (float32, error)
	GetUserFinancialReport(id uuid.UUID, period *domain.Period) (*domain.FinancialReportByPeriod, error)
}
