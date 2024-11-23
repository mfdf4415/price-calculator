package prices

import (
	"fmt"
	"price-calculator/convertion"
	filemanager "price-calculator/fileManager"
)

type TaxIncludePricesJob struct {
	TaxRate          float64
	Prices           []float64
	TaxIncludePrices map[string]string
}

func (job *TaxIncludePricesJob) LoadedData() {

	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := convertion.StrToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices = prices
}

func (job TaxIncludePricesJob) Process() {
	job.LoadedData()
	result := make(map[string]string)

	for _, price := range job.Prices {
		priceIncludeTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%2.f", price)] = fmt.Sprintf("%2.f", priceIncludeTax)
	}

	job.TaxIncludePrices = result
	filemanager.WriteJSONToFile(fmt.Sprintf("prices_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludePricesJob(taxRate float64) *TaxIncludePricesJob {
	return &TaxIncludePricesJob{
		TaxRate: taxRate,
		Prices:  []float64{10, 20, 30},
	}
}
