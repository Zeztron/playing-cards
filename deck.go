package main

import "fmt"

// Create a new type of deck
// Which is a type of strings
type deck []string

// The d is a reference to cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
