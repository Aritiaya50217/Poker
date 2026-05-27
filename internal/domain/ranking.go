package domain

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeKind
	Straight
	Flush
	FullHouse
	FourKind
	StraightFlush
	RoyalFlush
)

var RankName = map[int]string{

	HighCard:      "High Card",
	OnePair:       "One Pair",
	TwoPair:       "Two Pair",
	ThreeKind:     "Three of Kind",
	Straight:      "Straight",
	Flush:         "Flush",
	FullHouse:     "Full House",
	FourKind:      "Four of Kind",
	StraightFlush: "Straight Flush",
	RoyalFlush:    "Royal Flush",
}
