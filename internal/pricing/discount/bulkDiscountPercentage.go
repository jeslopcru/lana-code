package discount

// A price rule for bulk discount: buying x or more of a product, the price of that product is reduced.
type BulkDiscountPercentage struct {
	DiscountPercentage float64
	MinUnits           int
}

func (discount *BulkDiscountPercentage) Calculate(quantity int, unitPrice float64) float64 {
	if quantity < discount.MinUnits {
		return unitPrice * float64(quantity)
	}
	return unitPrice * float64(quantity) * (100 - discount.DiscountPercentage) / 100
}
