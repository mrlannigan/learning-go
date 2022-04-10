package main

// "math/rand"

func main() {
	// cards := newDeck()

	// savedResult := cards.saveToFile("./test.txt")

	// if savedResult == nil {
	// 	fmt.Println("Wrote file successfully")
	// } else {
	// 	fmt.Println("Error occurred:", savedResult)
	// }

	// rand.Seed(time.Now().UnixNano())

	cards := newDeckFromFile("./test.txt")
	cards.shuffle()
	cards.print()
}
