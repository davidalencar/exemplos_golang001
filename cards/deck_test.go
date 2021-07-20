package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16, but got %v", len(d))
	}

	if d[0] != "As de Espadas" {
		t.Errorf("Expected first card of As de Espadas, but got %v", d[0])
	}

}
