package converter

func UsdToIdr(amount float64) float64 {
	rate := 15600.0
	return amount * rate
}