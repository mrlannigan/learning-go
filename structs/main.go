package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jimmy := person{
		firstName: "Jimmy",
		lastName:  "Toledo",
		contact: contactInfo{
			email:   "jimmy@trucks.com",
			zipCode: 43056,
		},
	}

	fmt.Printf("%+v", jimmy)
}
