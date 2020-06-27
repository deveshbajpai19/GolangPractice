package main

import "fmt"

func main()  {

	//slice of string
	cards := newDeck()
	cards.print()
	fmt.Println("----------------------")
	hand, remainingCards := deal(cards, 2)
	hand.print()
	fmt.Println("----------------------")
	remainingCards.print()

}

