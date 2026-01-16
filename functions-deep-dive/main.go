package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(0, numbers...)
	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, val := range numbers {
		sum += val
	}
	return sum
}
