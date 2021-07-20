package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16, but got %v", len(d))
	}

	if d[0] != "As de Espadas" {
		t.Errorf("Expected first card of As de Espadas, but got %v", d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedD := newDeckFromFile("_decktesting")

	if len(loadedD) != 16 {
		t.Errorf("Expected 16, got %v", len(loadedD))
	}
	os.Remove("_decktesting")
}
