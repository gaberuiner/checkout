package clientcreateorder

import (
	"context"

	domain "route256/checkout/internal/domain/purchasedomain"
)

func (c *Client) ListCart(ctx context.Context, request domain.ListCartRequest) (domain.ListCartResponse, error) {
	return domain.ListCartResponse{}, nil
}
