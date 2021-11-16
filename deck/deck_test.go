package deck

import (
	"github.com/adrianprayoga/gophercises/deck/card"
	"testing"
)

func TestAbsSort(t *testing.T) {
	deck := New(AddJoker(4))

	var jokerCount int
	for _, c := range deck {
		if c.Face == card.Joker {
			jokerCount++
		}
	}

	if jokerCount != 4 {
		t.Errorf("got %v, want %v", jokerCount, 4)
	}
}
