package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	var appUser user
	appUser = user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
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
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
