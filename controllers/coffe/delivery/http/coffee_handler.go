package http

import (
	"github.com/jfalih/goku/domain"
	"github.com/julienschmidt/httprouter"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CoffeeHandler struct {
	CUsecase domain.CoffeeUseCase
}

func NewCoffeeHandler(router *httprouter.Router, us domain.Coffee) {
	handler := &CoffeeHandler{
		CUsecase: us,
	}
	router.GET("/coffee", handler.FetchCoffee)
}

func (a *CoffeeHandler) FetchArticle() error {
}
