package pricing

import (
	"fmt"
	"lana-challenge/internal/pricing/discount"
)

type PriceRules struct {
	products  map[string]Product
	discounts map[string]discount.Discount
}

func NewPriceRules() *PriceRules {
	calculator := &PriceRules{}
	calculator.products = make(map[string]Product)
	calculator.discounts = make(map[string]discount.Discount)
	return calculator
}

func (pr *PriceRules) AddProduct(product Product) error {
	_, ok := pr.products[product.Code]
	if !ok {
		pr.products[product.Code] = product
	}
	return nil
}

func (pr *PriceRules) SetDiscount(code string, discount discount.Discount) error {
	_, ok := pr.products[code]
	if !ok {
		return fmt.Errorf("product %s doesn't exists", code)
	}
	_, ok = pr.discounts[code]
	if ok {
		return fmt.Errorf("product %s have already discount", code)
	}
	pr.discounts[code] = discount
	return nil
}
