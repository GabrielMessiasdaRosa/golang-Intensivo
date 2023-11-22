package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsEmpty(test *testing.T) {
	order := Order{}
	assert.Error(test, order.Validate(), "ID is empty")
}

func TestIfItGetsAnErrorIfPriceIsInvalid(test *testing.T) {
	order := Order{
		ID: "123",
		Price: -1,
	}
	assert.Error(test, order.Validate(), "Price is invalid")
}

func TestIfItGetsAnErrorIfTaxIsInvalid(test *testing.T) {
	order := Order{
		ID: "123",
		Price: 1,
		Tax: -1,
	}
	assert.Error(test, order.Validate(), "Tax is invalid")
}

func TestIfItGetsAnErrorIfFinalPriceIsInvalid(test *testing.T) {
	order := Order{
		ID: "123",
		Price: -1,
		Tax: -1,
	}
	assert.Error(test, order.CalculateFinalPrice(), "Final price is invalid")
}

func TestIfItGetsAnErrorIfOrderIsValid(test *testing.T) {
	order := Order{
		ID: "123",
		Price: 1,
		Tax: 1,
	}
	assert.NoError(test, order.Validate())
}

func TestIfItCreatesAnOrder(test *testing.T) {
	order, err := NewOrder("123", 1, 1)
	assert.NoError(test, err)
	assert.Equal(test, "123", order.ID)
	assert.Equal(test, 1.0, order.Price)
	assert.Equal(test, 1.0, order.Tax)
	assert.Equal(test, 2.0, order.FinalPrice)
}
