package domain

import "errors"

const MaxCards = 5

type Hand struct {
	cards []Card
}

func NewHand() Hand {
	return Hand{cards: make([]Card, 0, MaxCards)}
}

func (h *Hand) Add(card Card) error {
	if len(h.cards) >= MaxCards {
		return errors.New("hand full. ")
	}

	h.cards = append(h.cards, card)
	return nil
}

func (h Hand) Cards() []Card {

	result := make([]Card, len(h.cards))

	copy(result, h.cards)

	return result
}
