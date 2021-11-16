package deck

import (
	"github.com/adrianprayoga/gophercises/deck/card"
	"math/rand"
	"sort"
)

type Deck struct {
	cards []card.Card
}

type Options func(*[]card.Card)

func New(opts ...Options) []card.Card {
	cards := createBase52Deck()

	for _, option := range opts {
		option(&cards)
	}

	return cards
}

func createBase52Deck() []card.Card {
	faceList := []int{card.Spades, card.Heart, card.Club, card.Diamond}
	numberList := makeRange(1, 13)

	cards := make([]card.Card, 52)

	loc := 0
	for _, f := range faceList {
		for _, n := range numberList {
			cards[loc] = card.Card{Face: f, Number: n}
			loc++
		}
	}

	return cards
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func SortDeck(less func(cs *[]card.Card) func(i, j int) bool) Options {
	return func(cards *[]card.Card) {
		sort.Slice(*cards, less(cards))
	}
}

func AddJoker(n int) Options {
	return func(cards *[]card.Card) {
		for i := 0; i < n; i++ {
			*cards = append(*cards, card.Card{card.Joker, 0})
		}
	}
}

func Shuffle() Options {
	return func(cards *[]card.Card) {
		cardLoc := *cards
		for i := range *cards {
			j := rand.Intn(i + 1)
			cardLoc[i], cardLoc[j] = cardLoc[j], cardLoc[i]
		}
	}
}

func FilterCards(isRemoved func(cs card.Card) bool) Options {
	return func(cards *[]card.Card) {
		for i, card := range *cards {
			if isRemoved(card) {
				*cards = remove(*cards, i)
			}
		}
	}
}

func AppendNewDeck(i int) Options {
	return func(cards *[]card.Card) {
		for j := 0; j < i; j++ {
			*cards = append(*cards, createBase52Deck()...)
		}
	}
}

func remove(cards []card.Card, i int) []card.Card {
	cards[i] = cards[len(cards)-1]
	return cards[:len(cards)-1]
}

func AscDeckSort(cs *[]card.Card) func(i int, j int) bool {
	return func(i int, j int) bool {
		return card.IsAbsLess((*cs)[i], (*cs)[j])
	}
}

func DescDeckSort(cs *[]card.Card) func(i int, j int) bool {
	return func(i int, j int) bool {
		return !card.IsAbsLess((*cs)[i], (*cs)[j])
	}
}
