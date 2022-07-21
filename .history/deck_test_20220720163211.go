package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 26 {
		t.Errorf("Expected deck length of 16 but, but got %v", len(d))
	}
}
