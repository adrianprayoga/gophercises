package card

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAbsSort(t *testing.T) {
	d10 := Card{Diamond, 10}
	s1 := Card{Spades, 1}
	h8 := Card{Heart, 8}
	h2 := Card{Heart, 2}

	var tests = []struct {
		cards    []Card
		expected bool
	}{
		{[]Card{d10, s1}, true},
		{[]Card{h8, h2}, false},
		{[]Card{h2, h8}, true},
		{[]Card{s1, h8}, false},
	}

	for i, tt := range tests {
		t.Run("comparison "+fmt.Sprint(i), func(t *testing.T) {
			ans := IsAbsLess(tt.cards[0], tt.cards[1])
			if !reflect.DeepEqual(ans, tt.expected) {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
