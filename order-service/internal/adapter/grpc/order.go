package grpchandler

import (
	"context"
	"order-service/internal/domain"
	"order-service/internal/usecase"
	orderproto "order-service/proto/order"
)

type OrderServer struct {
	orderproto.UnimplementedOrderServiceServer
	orderUseCase usecase.OrderUseCase
}

func NewOrderServer(orderUseCase usecase.OrderUseCase) *OrderServer {
	return &OrderServer{orderUseCase: orderUseCase}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderproto.CreateOrderRequest) (*orderproto.OrderResponse, error) {
	order := domain.Order{
		UserID:        req.UserId,
		Status:        req.Status,
		PaymentStatus: req.PaymentStatus,
		TotalAmount:   req.TotalAmount,
	}

	id, err := s.orderUseCase.CreateOrder(ctx, order, mapOrderItemsFromProto(req.Items))
	if err != nil {
		return nil, err
	}

	return &orderproto.OrderResponse{
		Id:            int32(id),
		UserId:        req.UserId,
		Status:        req.Status,
		PaymentStatus: req.PaymentStatus,
		TotalAmount:   req.TotalAmount,
	}, nil
}

func (s *OrderServer) GetOrderByID(ctx context.Context, req *orderproto.GetOrderRequest) (*orderproto.OrderResponse, error) {
	order, orderItems, err := s.orderUseCase.GetOrderByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &orderproto.OrderResponse{
		Id:            int32(order.ID),
		UserId:        int32(order.UserID),
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt.String(),
		UpdatedAt:     order.UpdatedAt.String(),
		Items:         mapOrderItemsToProto(orderItems),
	}, nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderproto.UpdateOrderStatusRequest) (*orderproto.OrderResponse, error) {
	err := s.orderUseCase.UpdateOrderStatus(ctx, req.Id, req.Status, req.PaymentStatus)
	if err != nil {
		return nil, err
	}

	order, orderItems, err := s.orderUseCase.GetOrderByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &orderproto.OrderResponse{
		Id:            int32(order.ID),
		UserId:        int32(order.UserID),
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt.String(),
		UpdatedAt:     order.UpdatedAt.String(),
		Items:         mapOrderItemsToProto(orderItems),
	}, nil
}

func (s *OrderServer) ListUserOrders(ctx context.Context, req *orderproto.ListOrdersRequest) (*orderproto.ListOrdersResponse, error) {
	orders, err := s.orderUseCase.ListOrders(ctx, nil, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	var orderResponses []*orderproto.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, &orderproto.OrderResponse{
			Id:            int32(order.ID),
			UserId:        int32(order.UserID),
			Status:        order.Status,
			PaymentStatus: order.PaymentStatus,
			TotalAmount:   order.TotalAmount,
			CreatedAt:     order.CreatedAt.String(),
			UpdatedAt:     order.UpdatedAt.String(),
		})
	}

	return &orderproto.ListOrdersResponse{
		Orders: orderResponses,
	}, nil
}

func mapOrderItemsFromProto(items []*orderproto.OrderItemRequest) []domain.OrderItem {
	var orderItems []domain.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductID:       item.ProductId,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}
	return orderItems
}

func mapOrderItemsToProto(items []domain.OrderItem) []*orderproto.OrderItemResponse {
	var orderItems []*orderproto.OrderItemResponse
	for _, item := range items {
		orderItems = append(orderItems, &orderproto.OrderItemResponse{
			ProductId:       int32(item.ProductID),
			Quantity:        int32(item.Quantity),
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}
	return orderItems
}
