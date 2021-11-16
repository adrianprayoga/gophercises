package card

const (
	Diamond = iota
	Club    = iota
	Heart   = iota
	Spades  = iota
	Joker   = iota
)

type Card struct {
	Face   int
	Number int
}

func (card Card) IsValid() bool {
	if card.Face != Spades && card.Face != Heart && card.Face != Club && card.Face != Diamond {
		return false
	}

	if card.Number < 1 || card.Number > 13 {
		return false
	}

	return true
}

func IsAbsLess(card1, card2 Card) bool {
	if card1.Face == Joker {
		return false
	}

	if card1.Face == card2.Face {
		return card1.Number < card2.Number
	}

	return card1.Face < card2.Face
}
