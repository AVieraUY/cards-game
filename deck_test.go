package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Want: 16, got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Want: Ace of Spades, got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Want: Four of Clubs, got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Want: 16, got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
