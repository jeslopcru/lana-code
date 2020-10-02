package storage

import (
	"lana-challenge/internal/pricing"
	"lana-challenge/internal/pricing/discount"
)

var Baskets = make(map[string]pricing.Basket)

func NewInMemoryRepository() *map[string]pricing.Basket {
	return &Baskets
}

func InitialData() *pricing.PriceRules {
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
