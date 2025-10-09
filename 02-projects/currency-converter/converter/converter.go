package converter

import "fmt"

type CurrencyType int

// enum constant
const (
	USD CurrencyType = iota // 0
	EUR                     // 2
	JPY                     // 3
	GBP                     // 4
	CAD                     // 5
)

// store conversion rates
// Rates relative to the Base Currency (e.g., 1 USD = X of the other currency)
var rates = map[CurrencyType]float64{
	USD: 1.0,  // Base rate
	EUR: 0.92, // Example: 1 USD = 0.92 EUR
	JPY: 145.0,
	GBP: 0.80,
	CAD: 1.35,
}

type Float interface {
	float32 | float64
}

// main currency converter
func Convert[T Float](amount T, from CurrencyType, to CurrencyType) (T, error) {
	fromRate, okFrom := rates[from]
	toRate, okTo := rates[to]

	if !okFrom {
		return 0, fmt.Errorf("unsupported source currency: %d", from)
	}
	if !okTo {
		return 0, fmt.Errorf("unsupported source currency: %d", to)
	}

	// conversion logic
	// convert amount to base currency
	amountInUSD := float64(amount) / fromRate

	// conver from base 'to' currency
	resultFloat64 := amountInUSD * toRate

	return T(resultFloat64), nil
}
