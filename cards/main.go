package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("My Hand:")
	hand.print()

	fmt.Println("\nRemaining Cards:")
	remainingDeck.print()
}
