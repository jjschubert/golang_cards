package main

import "fmt"

//create a new type of 'deck'
//which is a slice of type string
type deck []string

//anytime someone calls new deck, will return a value of type deck
//that's what the deck is doing there
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//replaced i and j with underscores, because otherwise go requires we return them
	//an underscore tells go these are temporary/not returned
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) is an example of a receiver
//gives any variable of type deck gets access to the print method (this function)
//custom to use first letter of type as variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//call deal with value type deck and value type int
//cannot be called with other types
//again, d is an instance of type deck
//(deck, deck) days that we will return two values, both type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
