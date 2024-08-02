package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) OutPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthday are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email,
		password,
		User{
			firstName: "ADMIN",
			lastName:  "LAST",
			birthdate: "Birthday",
			createdAt: time.Now(),
		},
	}
}
