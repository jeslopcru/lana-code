package lana_challenge

import (
	"github.com/stretchr/testify/assert"
	"lana-challenge/internal/pricing"
	"lana-challenge/internal/pricing/discount"
	"testing"
)

func createPriceRules() *pricing.PriceRules {
	priceRules := pricing.NewPriceRules()
	pen := pricing.Product{"PEN", "Lana Pen", 5.0}
	_ = priceRules.AddProduct(pen)
	_ = priceRules.SetDiscount("PEN", &discount.BuyXGetY{2, 1})
	mug := pricing.Product{"MUG", "Lana Coffee Mug", 7.5}
	_ = priceRules.AddProduct(mug)

	tshirt := pricing.Product{"TSHIRT", "Lana T-Shirt", 20.0}
	_ = priceRules.AddProduct(tshirt)
	_ = priceRules.SetDiscount("TSHIRT", &discount.BulkDiscountPercentage{25.0, 3})

	return priceRules
}

var basketCases = []struct {
	codes    []string
	expected float64
}{
	{codes: []string{"PEN", "TSHIRT", "MUG"}, expected: 32.50},
	{codes: []string{"PEN", "TSHIRT", "PEN"}, expected: 25.0},
	{codes: []string{"TSHIRT", "TSHIRT", "TSHIRT", "PEN", "TSHIRT"}, expected: 65.0},
	{codes: []string{"PEN", "TSHIRT", "PEN", "PEN", "MUG", "TSHIRT", "TSHIRT"}, expected: 62.5},
}

func TestBasketTotal(t *testing.T) {
	for _, bc := range basketCases {
		checkout := pricing.NewBasket(createPriceRules())
		for _, code := range bc.codes {
			_ = checkout.AddProduct(code)
		}
		assert.Equal(t, checkout.Total(), bc.expected)
	}
}
