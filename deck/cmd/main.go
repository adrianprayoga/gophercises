package main

import (
	"deck"
	"deck/card"
	"fmt"
)

func main() {
	cards := deck.New(
		deck.AddJoker(3),
		deck.AppendNewDeck(1),
		deck.FilterCards(func(cs card.Card) bool {
			return cs.Face == card.Spades
		}),
	)
	for _, c := range cards {
		fmt.Println(c)
	}
}
