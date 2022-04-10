package main

import "fmt"

func main() {
	nums := make([]int, 11)

	for i := range nums {
		state := "odd"
		if i%2 == 0 {
			state = "even"
		}

		fmt.Printf("%v is %v\n", i, state)
	}
}
