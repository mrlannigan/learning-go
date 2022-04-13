package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jimmy := person{
		firstName: "Jimmy",
		lastName:  "Toledo",
		contactInfo: contactInfo{
			email:   "jimmy@trucks.com",
			zipCode: 43056,
		},
	}

	jimmy.print()
	jimmy.updateName("")
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newName string) {
	splitName := strings.Split(newName, " ")
	p.firstName = splitName[0]
	p.lastName = strings.Join(splitName[1:], " ")
}
