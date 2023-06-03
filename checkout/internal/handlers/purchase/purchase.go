package purchase

import (
	"context"
	"errors"
	"log"

	"route256/checkout/internal/domain/purchasedomain"
)

var ErrUserNotFound = errors.New("user not found")

type Handler struct {
	Model *purchasedomain.Model
}

type Request struct {
	User int64 `json:"user"`
}

type Response struct {
	OrderID int64
}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("%+v", req)
	err := h.Model.Purchase(ctx, req.User)
	return Response{OrderID: 1}, err
}
