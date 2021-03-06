package repository

import (
	"context"

	"github.com/h-tachikawa/mechanical-receptionist/domain"
)

type VisitHistoryRepository interface {
	GetLatest(ctx context.Context) (*domain.VisitHistory, error)
	Add(ctx context.Context, doc *domain.VisitHistory) error
}
