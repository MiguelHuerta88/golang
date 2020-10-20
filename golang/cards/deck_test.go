package main

import (
	"os"
	"testing"
)

const DECKSIZE int = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// make sure that deck equals size
	if len(d) != DECKSIZE {
		var m = "Expected deck length of %d but got %d"
		t.Errorf(m, DECKSIZE, len(d))
	}

	// pull first card from deck and make sure that its equal to "Ace of Spades"
	fc := "Ace of Spades"
	fi := 0
	if d[fi] != fc {
		t.Errorf("Expected first card of %s, but got %s", fc, d[fi])
	}

	lc := "King of Clubs"
	li := len(d) - 1
	if d[li] != lc {
		t.Errorf("Expected first card of %s, but got %s", lc, d[li])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove testing file in case test crashed while running
	df := "_deckTesting"

	os.Remove(df)

	// create deck
	d := newDeck()
	d.saveToFile(df)

	loadedDeck := newDeckFromFile(df)

	// assert length
	if DECKSIZE != len(loadedDeck) {
		t.Errorf("Expected %d cards in deck, got %d", DECKSIZE, len(loadedDeck))
	}

	os.Remove(df)
}
