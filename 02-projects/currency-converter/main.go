package main

import (
	"fmt"

	"github.com/ShahandFahad/golang-sandbox/02-projects/currency-converter/converter"
)

func main() {
	fmt.Println("Welcome to Currency Converter")
	
	// enum constant form converter package
	amount := 100.0 // float64
	from := converter.EUR
	to := converter.JPY


	// Test 1
	// call the generic
	result, err := converter.Convert(amount, from, to)
	if err != nil {
		fmt.Println("Conversion error: ", err)
		return
	}

	fmt.Printf("%.2f %v is equal to %.2f %v\n", amount, from, result, to)

	// Test 2
	amountF32 := float32(50)
    resultF32, err := converter.Convert(amountF32, converter.GBP, converter.CAD)
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }
    fmt.Printf("%.2f %v is equal to %.2f %v\n", amountF32, converter.GBP, resultF32, converter.CAD)

}
