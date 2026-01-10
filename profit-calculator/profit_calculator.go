package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("Revenue: ")

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInput("Expenses: ")

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	taxRate := getUserInput("Tax Rate: ")

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate / 100)
	// ratio := earningsBeforeTax / profit
	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	
	fmt.Printf("%.2f\n", earningsBeforeTax)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}