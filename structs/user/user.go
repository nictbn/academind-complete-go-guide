package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	FirstName string
	LastName string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email string, password string) Admin {
	return Admin {
		email: email,
		password: password,
		User: User {
			FirstName: "ADMIN",
			LastName: "ADMIN",
			BirthDate: "___",
			CreatedAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func New(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name and Birth Date are required.")
	}
	return &User {
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}
