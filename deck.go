package main

import "fmt"

// Create a new type of deck
// Which is an array of strings
type deck []string

// Creates and returns a list of playing cards
// Returns a value of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Replaced iterators with underscores if iterators are not used.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// Append it to cards
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// The d is a reference to cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
