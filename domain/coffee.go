package domain

import "context"

type Coffee struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CoffeeUseCase interface {
	FetchCoffee(ctx context.Context) (res []Coffee, err error)
}
