package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("NewDeck returned wrong number of cards, expected 52 got %v", len(d))
	}
}
