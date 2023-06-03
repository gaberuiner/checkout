package clientcreateorder

import (
	"context"

	domain "route256/checkout/internal/domain/purchasedomain"
)

func (c *Client) CreateOrder(ctx context.Context, request domain.CreateOrderRequest) (domain.CreateOrderResponse, error) {
	return domain.CreateOrderResponse{}, nil
}
