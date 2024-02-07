package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected 16 but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")

	deck := newDeck()
	deck.saveToFile("_decktest")

	loaddeck := newDeckFromFile("_decktest")
	if len(loaddeck) != 16 {
		t.Errorf("error deck length %v", len(loaddeck))
	}

	os.Remove("_decktest")


}
