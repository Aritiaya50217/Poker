package domain

import "fmt"

type Suit string
type Rank int

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

type Card struct {
	rank Rank
	suit Suit
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank: rank, suit: suit}
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) String() string {
	return fmt.Sprintf("[%s %s]", rankToString(c.rank), c.suit)
}

func rankToString(r Rank) string {
	switch r {
	case 14:
		return "A"
	case 13:
		return "K"
	case 12:
		return "Q"
	case 11:
		return "J"
	default:
		return fmt.Sprintf("%d", r)
	}
}
