package convertion

import (
	"errors"
	"strconv"
)

func StrToFloat(strings []string) ([]float64, error) {

	convertedSlice := make([]float64, len(strings))

	for index, value := range strings {
		floated, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, errors.New("list convertion faild")
		}

		convertedSlice[index] = floated
	}

	return convertedSlice, nil
}
