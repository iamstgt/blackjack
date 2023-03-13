package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
	}
}
