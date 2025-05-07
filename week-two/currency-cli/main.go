package main

import (
	"fmt"
	"currency-cli/converter"
)

func main() {
	amount := 100.0
    converted := converter.UsdToIdr(amount)
    fmt.Printf("$%.2f = Rp%.2f\n", amount, converted)
}