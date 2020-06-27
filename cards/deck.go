package main

import "fmt"

//create a type of 'deck' which is a slice of strings
type deck []string


func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, cardSuit := range cardSuits {

		for _, cardValue := range cardValues {

			cards = append(cards, cardValue + " of " + cardSuit)
		}

	}

	return cards
}

//receiver function
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}
