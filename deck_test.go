package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected lenght of 16 , but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected ace of spade , but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Club , but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards from file , got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
