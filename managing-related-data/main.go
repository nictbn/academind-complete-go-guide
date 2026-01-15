package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames[1] = "Mike"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)
}
