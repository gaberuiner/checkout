package purchasedomain

import (
	"context"
	"fmt"

	"route256/libs/clnwrapper"
)

type CreateOrder interface {
	CreateOrder(ctx context.Context, request CreateOrderRequest) (CreateOrderResponse, error)
	ListCart(ctx context.Context, request ListCartRequest) (ListCartResponse, error)
}
type ListCartRequest struct {
	User int64 `json:"user"`
}
type ListCartResponse struct {
	Items []Item `json:"items"`
}

type CreateOrderRequest struct {
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

type Model struct {
	createOrder CreateOrder
}

func New(createOrder CreateOrder) *Model {
	return &Model{createOrder: createOrder}
}

func (m *Model) Purchase(ctx context.Context, user int64) error {
	path := "http://localhost:8080/listCart"
	req := ListCartRequest{User: user}
	clientListCart := clnwrapper.New(ctx, path, req, m.createOrder.ListCart)
	ListRes, err := clientListCart.Wrap()
	if err != nil {
		return fmt.Errorf("get listcart: %w", err)
	}
	path = "http://localhost:8081/createOrder"
	clientCreateOrder := clnwrapper.New(ctx, path, CreateOrderRequest{User: user, Items: ListRes.Items}, m.createOrder.CreateOrder)

	OrderRes, err := clientCreateOrder.Wrap()
	if err != nil {
		return fmt.Errorf("make order: %w", err)
	}
	if OrderRes.OrderID == 0 {
		return fmt.Errorf("didnt make order")
	}
	return nil
}
