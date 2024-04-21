package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := getInputFloat64("Investment Amount: ")
	expectedReturnRate := getInputFloat64("Expected Return Rate: ")
	years := getInputFloat64("Years: ")
	inflationRate := getInputFloat64("Expected Average Inflation Rate: ")
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func getInputFloat64(message string) (returnValue float64) {
	fmt.Print(message)
	fmt.Scan(&returnValue)

	return returnValue
}

func calculateFutureValues(invAmt, expRtnRate, years, infRate float64) (fv float64, rfv float64) {
	fv = invAmt * math.Pow(1+expRtnRate/100, years)
	rfv = fv / math.Pow(1+infRate/100, years)

	return fv, rfv
}
