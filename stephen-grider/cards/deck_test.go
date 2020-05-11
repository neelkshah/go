package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length equal to 52, got length %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Two of Clubs" {
		t.Errorf("Expected Two of Clubs, got %v", d[len(d)-1])
	}
}
