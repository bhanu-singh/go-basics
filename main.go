package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	//cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//fmt.Println(remainingCards.toString())

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))
	// card2 := newDeckFromFile("my_cards")
	// card2.print()
}
