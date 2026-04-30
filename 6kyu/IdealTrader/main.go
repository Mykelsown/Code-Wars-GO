package main

import "fmt"

func main() {
	test := []float64{2.0, 2.0}
	fmt.Println(idealTrader(test))
}

func idealTrader(prices []float64) float64 {
	amountRatio := 1.0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			amountRatio = (amountRatio * prices[i+1]) / prices[i]
		}
	}
	return amountRatio
}
