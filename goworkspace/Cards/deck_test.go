package main

import "testing"

func TestNewDeck(t *testing.T) {

	cards := getNewDeck()

	expected := 16
	actual := len(cards)

	if actual != expected {
		t.Errorf("Expected deck length of 52, but got %v", actual)
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", cards[len(cards)-1])
	}
}
