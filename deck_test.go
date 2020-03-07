package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Seven of Clubs" {
		t.Errorf("Expected last card of 'Seven of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveFromDeckAndNewDeckFromFile(t *testing.T) {
	// delete testing file at the start and end of function run
	fn := "_decktesting.txt"
	os.Remove(fn)

	deck := newDeck()
	deck.saveToFile(fn)

	loadedDeck := newDeckFromFile(fn)

	if len(loadedDeck) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(loadedDeck))
	}

	os.Remove(fn)
}
