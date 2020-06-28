package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := newDeck()
	if len(cards) != 56 {
		t.Errorf("Expected length of deck as 56 but got %d", len(cards))
	}

	if cards[0] != "One of Spades" {
		t.Errorf("Expected first card to be One of Spades but got %s", cards[0])
	}

	if cards[len(cards)-1] != "Ace of Hearts" {
		t.Errorf("Expected larst card to be Ace of Hearts but got %s", cards[len(cards)-1])
	}
}

func TestFileIoDeck(t *testing.T) {

	testFileName := "_deckTesting"

	//pre-test cleanup
	os.Remove(testFileName)

	cards := newDeck()
	cards.saveToFile(testFileName)
	loadedCards := newDeckFromFile(testFileName)
	if len(loadedCards) != 56 {
		t.Errorf("Expected length of deck as 56 but got %d", len(loadedCards))
	}

	//post-test cleanup
	os.Remove(testFileName)
}
