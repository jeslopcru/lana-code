package pricing

import (
	"fmt"
	"github.com/google/uuid"
)

type Basket struct {
	UUID       string
	priceRules *PriceRules
	Products   map[string]int
}

func NewBasket(rules *PriceRules) *Basket {
	c := &Basket{priceRules: rules}
	c.Products = make(map[string]int)
	c.UUID = uuid.New().String()
	return c
}

func (c *Basket) AddProduct(productCode string) error {
	_, ok := c.priceRules.products[productCode]
	if !ok {
		return fmt.Errorf("product %s doesn't exists", productCode)
	}
	c.Products[productCode]++
	return nil
}

func (c *Basket) RemoveProducts() {
	c.Products = make(map[string]int)
}

func (c *Basket) Total() float64 {
	total := 0.0
	for productCode, quantity := range c.Products {
		product, _ := c.priceRules.products[productCode]
		totalByProduct := c.totalByProduct(productCode, product.Price, quantity)
		total += totalByProduct
	}
	return total
}

func (c *Basket) totalByProduct(code string, price float64, quantity int) float64 {
	discount, ok := c.priceRules.discounts[code]
	if ok {
		return discount.Calculate(quantity, price)
	}
	return float64(quantity) * price
}
