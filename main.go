package main

import (
	"price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.9, 0.1, 0.4}

	for _, v := range taxRates {

		jobPrices := prices.NewTaxIncludePricesJob(v)
		jobPrices.Process()
	}

}
