package usecase

import (
	"context"

	"github.com/IrentyM/store/statistics-service/internal/domain/model"
)

type StatsUseCase interface {
	HandleOrderEvent(ctx context.Context, event *model.OrderEvent) error
	HandleInventoryEvent(ctx context.Context, event *model.InventoryEvent) error
	GetUserOrderStatistics(ctx context.Context, userID int64) (*model.UserStatistics, error)
	GetUserHourlyStatistics(ctx context.Context, userID int64) (map[int]int, error)
}

type statsUseCase struct {
	repo StatsRepo
}

func NewStatsUseCase(repo StatsRepo) StatsUseCase {
	return &statsUseCase{repo: repo}
}

func (uc *statsUseCase) HandleOrderEvent(ctx context.Context, event *model.OrderEvent) error {
	return uc.repo.SaveOrderEvent(ctx, event)
}

func (uc *statsUseCase) HandleInventoryEvent(ctx context.Context, event *model.InventoryEvent) error {
	return uc.repo.SaveInventoryEvent(ctx, event)
}

func (uc *statsUseCase) GetUserOrderStatistics(ctx context.Context, userID int64) (*model.UserStatistics, error) {
	return uc.repo.GetUserOrderStatistics(ctx, userID)
}

func (uc *statsUseCase) GetUserHourlyStatistics(ctx context.Context, userID int64) (map[int]int, error) {
	events, err := uc.repo.GetUserOrderStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	hourlyStats := make(map[int]int)
	for _, event := range events {
		if event.Status == "completed" {
			hour := event.CreatedAt.Hour()
			hourlyStats[hour]++
		}
	}
	return hourlyStats, nil
}
