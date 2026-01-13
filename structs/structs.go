package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if you maintain the order, you can also instantiate like this, regardless of the variable names:
	// appUser = user{
	// 	firstName,
	// 	lastName,
	// 	birthDate,
	// 	time.Now(),
	// }

	// you can also instantiate a null user
	// appUser = {}

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
