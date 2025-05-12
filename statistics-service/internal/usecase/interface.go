package usecase

import (
	"context"

	"github.com/IrentyM/store/statistics-service/internal/domain/model"
)

type StatsRepo interface {
	SaveOrderEvent(ctx context.Context, event *model.OrderEvent) error
	SaveInventoryEvent(ctx context.Context, event *model.InventoryEvent) error
	GetUserOrderStats(ctx context.Context, userID int64) ([]model.OrderEvent, error)
	GetUserOrderStatistics(ctx context.Context, userID int64) (*model.UserStatistics, error)
}
