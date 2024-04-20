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

	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	fmt.Print("Profit Before Tax: ")
	fmt.Println(ebt)

	fmt.Print("Profit After Tax: ")
	fmt.Println(profit)

	fmt.Print("Profit Ratio: ")
	fmt.Println(ratio)

}
