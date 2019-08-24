package main

func main() {
	// slice of string (array)
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// Notifying that the function will return a type of string
func newCard() string {
	return "Five of Diamonds"
}
