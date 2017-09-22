package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}

	if d[0] != "Ace of Wands" {
		t.Errorf("Expected first card of Ace of Wands, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Pentacles" {
		t.Errorf("Expected last card of King of Pentacles, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cards in deck, got %v", len(loadedDeck))

	}

	os.Remove("_decktesting")
}
