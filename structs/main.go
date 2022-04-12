package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	dude := person{lastName: "With A Nice Carpet"}

	test := fmt.Sprintf("%+v", dude)
	fmt.Println(test)
}
