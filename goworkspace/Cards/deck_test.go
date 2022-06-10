package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := getNewDeck()

	expected := 16
	actual := len(cards)

	if actual != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, actual)
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndGetNewDeckFromFile(t *testing.T) {
	filename := "./Data/_testdeskofcards.txt"

	os.Remove(filename)

	cards := getNewDeck()
	cards.saveToFile(filename)

	loadedDeck := getNewDeckFromFile(filename)

	expected := 16
	actual := len(loadedDeck)

	if actual != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, actual)
	}

	loadedDeck.print("Deck from File")

	os.Remove(filename)
}
