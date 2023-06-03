package clientstocks

import (
	"context"

	domain "route256/checkout/internal/domain/addtocartdomain"
)

func (c *Client) Stoks(ctx context.Context, requestStocks domain.StocksRequest) (domain.StocksResponse, error) {
	return domain.StocksResponse{}, nil
}
