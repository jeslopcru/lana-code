package discount

type BuyXGetY struct {
	Pay  int
	Free int
}

func (buyGet *BuyXGetY) Calculate(quantity int, price float64) float64 {
	free := (quantity / buyGet.Pay) * buyGet.Free
	remaining := quantity % buyGet.Pay
	return float64(free+remaining) * price

}
