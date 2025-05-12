package frontend

import (
	"context"
	"fmt"

	pb "github.com/IrentyM/store/apis/gen/statistics-service/service/frontend/statistics/v1"
	"github.com/IrentyM/store/statistics-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	pb.UnimplementedStatisticsServiceServer
	uc usecase.StatsUseCase
}

func NewGRPCHandler(uc usecase.StatsUseCase) *GRPCHandler {
	return &GRPCHandler{uc: uc}
}

func (h *GRPCHandler) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {

	stats, err := h.uc.GetUserOrderStatistics(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get order stats: %v", err)
	}

	hourly := make(map[string]int32)
	for hour, count := range stats.HourlyDistribution {
		hourly[formatHour(hour)] = int32(count)
	}

	return &pb.UserOrderStatisticsResponse{
		TotalOrders:        int32(stats.TotalOrders),
		HourlyDistribution: hourly,
	}, nil
}

func (h *GRPCHandler) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {

	stats, err := h.uc.GetUserOrderStatistics(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user stats: %v", err)
	}

	return &pb.UserStatisticsResponse{
		TotalItemsPurchased:  int32(stats.TotalItemsPurchased),
		AverageOrderValue:    stats.AverageOrderValue,
		MostPurchasedItem:    stats.MostPurchasedItem,
		TotalCompletedOrders: int32(stats.CompletedOrders),
	}, nil
}

func formatHour(hour int) string {
	return fmt.Sprintf("%02d:00", hour)
}
