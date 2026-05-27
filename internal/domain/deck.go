package domain

import (
	"errors"
	"time"

	"math/rand"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}

	cards := make([]Card, 0, 52)

	for _, s := range suits {
		for r := 2; r <= 14; r++ {
			cards = append(cards, NewCard(Rank(r), s))
		}
	}

	return &Deck{cards: cards}

}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck empty")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]

	return card, nil
}

func (d *Deck) Remaining() int {
	return len(d.cards)
}
