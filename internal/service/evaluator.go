package service

import (
	"sort"

	"github.com/Aritiaya50217/Poker/internal/domain"
)

type Result struct {
	Rank    int
	Kickers []int
}

func Evaluate(cards []domain.Card) Result {
	count := map[int]int{}
	suits := map[string]int{}

	var values []int

	for _, c := range cards {
		v := int(c.Rank())
		values = append(values, v)
		count[v]++

		suits[string(c.Suit())]++

	}

	sort.Ints(values)

	flush := len(suits) == 1

	straight := isStraight(values)

	if flush && straight && values[0] == 10 {
		return Result{Rank: domain.RoyalFlush, Kickers: []int{14}}
	}

	if flush && straight {
		return Result{Rank: domain.StraightFlush, Kickers: []int{values[4]}}
	}

	return evaluateKinds(count, values, flush, straight)
}

func evaluateKinds(count map[int]int, values []int, flush bool, straight bool) Result {
	var four int
	var three int
	pairs := []int{}

	for rank, c := range count {
		switch c {
		case 4:
			four = rank
		case 3:
			three = rank
		case 2:
			pairs = append(pairs, rank)
		}
	}

	sort.Ints(pairs)
	sort.Ints(values)

	// four of a kind
	if four > 0 {
		return Result{Rank: domain.FourKind, Kickers: []int{four}}
	}

	// full house
	if three > 0 && len(pairs) == 1 {
		return Result{
			Rank:    domain.FullHouse,
			Kickers: []int{three, pairs[0]},
		}
	}

	// flush
	if flush {
		return Result{
			Rank:    domain.Flush,
			Kickers: reverse(values),
		}
	}

	// straight
	if straight {
		return Result{Rank: domain.Straight, Kickers: []int{values[4]}}
	}

	// three of a kind
	if three > 0 {
		return Result{Rank: domain.ThreeKind, Kickers: []int{three}}
	}

	// two pair
	if len(pairs) == 2 {
		return Result{Rank: domain.TwoPair, Kickers: []int{
			pairs[1], // คู่ใหญ่
			pairs[0], // คู่เล็ก
		}}
	}

	// one pair
	if len(pairs) == 1 {
		return Result{Rank: domain.OnePair, Kickers: []int{pairs[0]}}
	}

	return Result{Rank: domain.HighCard, Kickers: reverse(values)}
}

func Compare(a, b Result) int {
	if a.Rank > b.Rank {
		return 1
	}

	if a.Rank < b.Rank {
		return -1
	}

	for i := 0; i < len(a.Kickers); i++ {
		if a.Kickers[i] > b.Kickers[i] {
			return -1
		}
	}
	return 0
}

func isStraight(v []int) bool {
	if len(v) != 5 {
		return false
	}

	if v[0] == 2 && v[1] == 3 && v[2] == 4 && v[3] == 5 && v[4] == 14 {
		return true
	}

	for i := 1; i < 5; i++ {
		if v[i] != v[i-1]+1 {
			return false
		}
	}

	return true
}

func reverse(v []int) []int {
	result := []int{}

	for i := len(v) - 1; i >= 0; i-- {
		result = append(result, v[i])
	}

	return result
}
