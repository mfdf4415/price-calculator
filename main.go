package main

import (
	"fmt"
	filemanager "price-calculator/fileManager"
	"price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.9, 0.1, 0.5}

	for _, v := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_%.0f.json", v*100))

		jobPrices := prices.NewTaxIncludePricesJob(v,fm)
		jobPrices.Process()
	}

}
