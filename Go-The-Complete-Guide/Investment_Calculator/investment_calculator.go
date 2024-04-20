package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount, expectedReturnRate, years, inflationRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Average Inflation Rate: ")
	fmt.Scan(&inflationRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value: %v\n", futureValue)
	fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
}
