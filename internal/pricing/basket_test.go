package pricing

import (
	"github.com/stretchr/testify/assert"
	"lana-challenge/internal/pricing/discount"
	"testing"
)

func TestWhenScanAProductThatNotExistThenObtainError(t *testing.T) {

	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)

	assert.Error(t, checkout.AddProduct("A_NOT_EXISTING_PRODUCT"))
}

func TestWhenScanAProductThatExistThenAddToBasket(t *testing.T) {
	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)

	assert.NoError(t, checkout.AddProduct("PEN"))
	assert.NotEmpty(t, checkout.Products)
}

func TestWhenRemoveTheBasketThenItIsEmpty(t *testing.T) {

	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)
	_ = checkout.AddProduct("PEN")
	assert.NotEmpty(t, checkout.Products)

	checkout.RemoveProducts()
	assert.Empty(t, checkout.Products)
}

func TestWhenAddProductWithANotDiscountedPriceThenTotalCalculatePrice(t *testing.T) {

	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)
	_ = checkout.AddProduct("MUG")

	total := checkout.Total()
	assert.Equal(t, 7.5, total)
}

func TestWhenTotalThenCalculatePriceApplyingBuyXGetYDiscount(t *testing.T) {

	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)
	_ = checkout.AddProduct("PEN")
	assert.NotEmpty(t, checkout.Products)

	total := checkout.Total()
	assert.Equal(t, 5.0, total)

	_ = checkout.AddProduct("PEN")

	totalWithDiscount := checkout.Total()
	assert.Equal(t, 5.0, totalWithDiscount)
}

func TestWhenTotalThenCalculatePriceApplyingBulkDiscountPercentage(t *testing.T) {

	priceRules := createPriceRulesWithADiscount()
	checkout := NewBasket(priceRules)
	_ = checkout.AddProduct("TSHIRT")

	total := checkout.Total()
	assert.Equal(t, 20.0, total)

	_ = checkout.AddProduct("TSHIRT")
	_ = checkout.AddProduct("TSHIRT")
	_ = checkout.AddProduct("TSHIRT")
	_ = checkout.AddProduct("PEN")

	totalWithDiscount := checkout.Total()
	assert.Equal(t, 65.0, totalWithDiscount)
}

func createPriceRulesWithADiscount() *PriceRules {
	priceRules := NewPriceRules()
	pen := Product{"PEN", "Lana Pen", 5.0}
	_ = priceRules.AddProduct(pen)
	_ = priceRules.SetDiscount("PEN", &discount.BuyXGetY{2, 1})
	mug := Product{"MUG", "Lana Coffee Mug", 7.5}
	_ = priceRules.AddProduct(mug)

	tshirt := Product{"TSHIRT", "Lana T-Shirt", 20.0}
	_ = priceRules.AddProduct(tshirt)
	_ = priceRules.SetDiscount("TSHIRT", &discount.BulkDiscountPercentage{25.0, 3})

	return priceRules
}
