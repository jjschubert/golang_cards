package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()

	//cards is our deck, 5 is the handSize we're passing
	//first value returned is hand, second is remainingCards
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
