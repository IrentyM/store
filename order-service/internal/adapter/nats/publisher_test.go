// order-service/internal/adapter/nats/publisher_test.go
package natsadapter

import (
	"context"
	"encoding/json"
	"order-service/internal/domain"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestOrderEventPublisher(t *testing.T) {
	// Mock NATS connection
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	publisher := NewOrderEventPublisher(nc)

	// Test subscription
	sub, err := nc.SubscribeSync("order.created")
	if err != nil {
		t.Fatalf("Failed to subscribe: %v", err)
	}

	// Test publish
	order := &domain.Order{
		ID:            1,
		UserID:        1,
		Status:        "created",
		PaymentStatus: "paid",
		TotalAmount:   100.0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	items := []domain.OrderItem{

		{
			ID:              1,
			OrderID:         1,
			ProductID:       1,
			Quantity:        2,
			PriceAtPurchase: 50.0,
		},
		{
			ID:              2,
			OrderID:         1,
			ProductID:       2,
			Quantity:        1,
			PriceAtPurchase: 50.0,
		},
	}

	err = publisher.PublishOrderCreated(context.Background(), order, items)
	if err != nil {
		t.Errorf("PublishOrderCreated failed: %v", err)
	}

	// Verify message
	msg, err := sub.NextMsg(1 * time.Second)
	if err != nil {
		t.Errorf("Failed to get message: %v", err)
	}

	var event OrderEventDTO
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		t.Errorf("Failed to unmarshal event: %v", err)
	}

	if event.OrderID != order.ID {
		t.Errorf("Expected order ID %d, got %d", order.ID, event.OrderID)
	}
}
