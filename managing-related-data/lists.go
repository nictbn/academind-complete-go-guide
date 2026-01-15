package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99, 12.99, 7.99)
	fmt.Println(updatedPrices, prices)
	prices = prices[1:]

	discountPrices := []float64{101.99, 80.99}
	prices = append(prices, discountPrices...)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[:3]
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[1:]
// 	fmt.Println(featuredPrices)

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)

// 	featuredPrices[0] = 199.99
// 	fmt.Println(prices)

// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
