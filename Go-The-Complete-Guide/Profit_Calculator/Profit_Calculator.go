/* Profit calculator
Asks for revenue, expenses, and tax rate input
calculates earnings before and after tax
ratio between the two values
outputs the three values
*/

package main

import (
	"fmt"
)

func main() {
	revenue := getInputFloat64("Revenue: ")
	expenses := getInputFloat64("Expenses: ")
	taxRate := getInputFloat64("Tax Rate: ")

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Printf("Profit Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}

func getInputFloat64(message string) (returnValue float64) {
	fmt.Print(message)
	fmt.Scan(&returnValue)

	return returnValue
}

func calculateProfits(rev, exp, tr float64) (ebt float64, pft float64, rt float64) {
	ebt = rev - exp
	pft = ebt * (1 - (tr / 100))
	rt = ebt / pft
	return ebt, pft, rt
}
