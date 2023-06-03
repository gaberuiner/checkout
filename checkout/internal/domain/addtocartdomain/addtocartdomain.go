package addtocartdomain

import (
	"context"
	"errors"
	"fmt"
	"log"

	"route256/libs/clnwrapper"
)

type StockChecker interface {
	Stoks(ctx context.Context, sku StocksRequest) (StocksResponse, error)
}

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksResponse struct {
	Stocks []struct {
		WarehouseID int64  `json:"warehouseID"`
		Count       uint64 `json:"count"`
	} `json:"stocks"`
}

type Model struct {
	stockChecker StockChecker
}

func New(stockChecker StockChecker) *Model {
	return &Model{stockChecker: stockChecker}
}

var ErrStockInsufficient = errors.New("stock insufficient")

func (m *Model) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	path := "http://localhost:8081/stocks"
	clientAddtocart := clnwrapper.New(ctx, path, StocksRequest{SKU: sku}, m.stockChecker.Stoks)
	stocks, err := clientAddtocart.Wrap()
	if err != nil {
		return fmt.Errorf("get stocks: %w", err)
	}
	log.Printf("stocks: %v", stocks)

	counter := int64(count)
	for _, stock := range stocks.Stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}
	return ErrStockInsufficient
}
