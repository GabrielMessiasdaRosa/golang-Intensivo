package usecase

import (
    "github.com/GabrielMessiasdaRosa/golang-intensivo/internal/entity")

type OrderInput struct {
	ID	   string
	Price  float64
	Tax   float64
}

type OrderOutput struct {
	ID	   string
	Price  float64
	Tax   float64
	FinalPrice float64
}

// SOLID - Dependency Inversion Principle
type CalculateFinalPrice struct {
	OrderReposiotry entity.OrderRepositoryInterface
}

func CalculateFinalPriceFactory(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderReposiotry: orderRepository,
	}
}

func (calculateFinalPrice *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = calculateFinalPrice.OrderReposiotry.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}