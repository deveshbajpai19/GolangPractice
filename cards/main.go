package main

func main() {

	//slice of string
	cards := newDeck()
	/*cards.print()
	fmt.Println("----------------------")
	hand, remainingCards := deal(cards, 2)
	hand.print()
	fmt.Println("----------------------")
	remainingCards.print()*/

	//cards.saveToFile("my_cards")

	//d := newDeckFromFile("my_cards")
	//d.print()

	cards.shuffle()
	cards.print()

}
