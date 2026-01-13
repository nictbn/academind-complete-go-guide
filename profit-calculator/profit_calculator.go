package main

import "fmt"
import "errors"
import "os"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate / 100)
	// ratio := earningsBeforeTax / profit
	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	
	fmt.Printf("%.2f\n", earningsBeforeTax)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return userInput, nil
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}