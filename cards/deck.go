package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, cardSuit := range cardSuits {

		for _, cardValue := range cardValues {

			cards = append(cards, cardValue+" of "+cardSuit)
		}

	}

	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {

	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	strSlice := strings.Split(string(byteSlice), ",")
	return deck(strSlice)
}

func (d deck) toString() string {

	strSlice := []string(d)
	return strings.Join(strSlice, ",")
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().Unix()) //using current time (Unix epoc seconds) as seed
	r := rand.New(src)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i] //swap
	}
}
