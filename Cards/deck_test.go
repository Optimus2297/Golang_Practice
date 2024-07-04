package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected the first card to be of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	nd := newDeckFromFile("_deckTesting")

	if len(nd) != 16 {
		t.Errorf("Expected deck length to be of 16, but got %v", len(nd))
	}

	os.Remove("_deckTesting")
}