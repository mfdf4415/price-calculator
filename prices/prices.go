package prices

import (
	"fmt"
	"price-calculator/convertion"
	filemanager "price-calculator/fileManager"
)

type TaxIncludePricesJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	Prices           []float64               `json:"prices"`
	TaxIncludePrices map[string]string       `json:"tax_included-prices"`
}

func (job *TaxIncludePricesJob) LoadedData() {

	lines, err := job.IOManager.ReadLines()
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
	job.IOManager.WriteJSONToFile(job)
}

func NewTaxIncludePricesJob(taxRate float64, fm filemanager.FileManager) *TaxIncludePricesJob {
	return &TaxIncludePricesJob{
		TaxRate:   taxRate,
		Prices:    []float64{10, 20, 30},
		IOManager: fm,
	}
}
