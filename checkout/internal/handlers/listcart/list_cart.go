package listcart

import (
	"context"
	"errors"
	"log"

	"route256/checkout/internal/handlers/listcart/swagger"
)

var ErrUserNotFound = errors.New("user not found")

type Handler struct{}

type Request struct {
	User int64 `json:"user"`
}

type Response struct {
	Items      []Items `json:"items"`
	TotalPrice uint32  `json:"totalPrice"`
}

type Items struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("%+v", req)

	var Total uint32
	ResponseList := Response{Items: []Items{{Sku: 773297411, Count: 1}, {Sku: 773297411, Count: 2}}, TotalPrice: 0}
	newItems := []Items{}
	newSwagger := swagger.NewProductService("Nagnm1rZz685OJgCYYHFIQyz")
	for _, resp := range ResponseList.Items {
		respProduct, err := newSwagger.GetItems(resp.Sku)
		newItem := Items{Sku: resp.Sku, Count: resp.Count}
		if err != nil {
			return Response{}, err
		}
		newItem.Name = respProduct.Name
		newItem.Price = respProduct.Price

		Total += newItem.Price * uint32(newItem.Count)
		newItems = append(newItems, newItem)

	}
	ResponseList.Items = newItems
	ResponseList.TotalPrice = Total
	return ResponseList, nil
}
