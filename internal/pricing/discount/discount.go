package discount

type Discount interface {
	Calculate(quantity int, price float64) float64
}
