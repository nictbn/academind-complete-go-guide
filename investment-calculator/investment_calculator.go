package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)

	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	// formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRealFutureValue := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Print(formattedFutureValue)
	// fmt.Print(formattedRealFutureValue)

	fmt.Printf(`Future Value: %.2f
Future Value (adjusted for inflation): %.2f
`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}