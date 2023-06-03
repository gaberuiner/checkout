package deletefromcart

import (
	"context"
	"errors"
	"log"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type Handler struct{}

type Request struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	// pusto
}

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("%+v", req)
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	select {
	case <-time.After(5 * time.Second):
	case <-ctx.Done():
		log.Println(ctx.Err())
		return Response{}, ctx.Err()
	}
	return Response{}, nil
}
