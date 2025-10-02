package main

import "testing"

// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) !=16 {
		t.Errorf("Expected deck length of 16, but got", len(d))
	}
}