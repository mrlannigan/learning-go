package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 24
	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be an Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Six of Clubs" {
		t.Errorf("Expected last card in deck to be a Six of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 24 {
		t.Errorf("Expected 24 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
