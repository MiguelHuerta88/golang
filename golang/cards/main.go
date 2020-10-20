package main

func main() {
	/*
		// syntax var varname typeForvar = value to assign
		var card string = "Ace of Spades"
		fmt.Println(card)

		// implicit declaration
		card1 := "Aces of Spades Again"

		fmt.Println(card1)*/

	// code that will use method
	//fmt.Println(newCard())

	// slices can grow and shring dynamically. As oppose to arrays that are fixed in length
	/*cards := newDeck()
	// use deal
	hand, remainingCards := deal(cards, 5)
	//hand, remainingCards := cards.dealUpdate(5) DOESNT WORK DONT USE

	hand.print()
	remainingCards.print()*/

	// section of code to save to file
	/*cards := newDeck()
	// try to save to HD
	cards.saveToFile("my_cards")*/

	// section of code to read from file
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
