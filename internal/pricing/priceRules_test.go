package pricing

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"lana-challenge/internal/pricing/discount"
	"testing"
)

func TestPricingRulesSetupInvalidProductRule(t *testing.T) {
	priceRules := NewPriceRules()
	pen := Product{"PEN", "Lana Pen", 5.0}
	_ = priceRules.AddProduct(pen)
	_ = priceRules.SetDiscount("PEN", &discount.BuyXGetY{2, 1})

	invalidProduct := "FAKE_PRODUCT"
	stError := priceRules.SetDiscount(invalidProduct, &discount.BuyXGetY{1, 3})
	assert.Equal(t, stError, errors.New(fmt.Sprintf("product %s doesn't exists", invalidProduct)))
}

func TestAddProduct(t *testing.T) {
	calculator := NewPriceRules()

	p := Product{"MUG", "Lana Coffee Mug", 7.50}

	assert.Empty(t, calculator.products)
	assert.NoError(t, calculator.AddProduct(p))
	assert.NotEmpty(t, calculator.products)
}

func TestSetDiscount(t *testing.T) {
	calculator := NewPriceRules()

	buyXGetY := discount.BuyXGetY{Pay: 3, Free: 2}
	assert.Error(t, calculator.SetDiscount("NOT_EXISTING_PRODUCT", &buyXGetY))

	p := Product{"MUG", "Lana Coffee Mug", 7.50}
	_ = calculator.AddProduct(p)

	assert.NoError(t, calculator.SetDiscount("MUG", &buyXGetY))
	assert.NotEmpty(t, calculator.discounts)

	assert.Error(t, calculator.SetDiscount("MUG", &buyXGetY))
}
