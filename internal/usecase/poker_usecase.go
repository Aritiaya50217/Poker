package usecase

import (
	"fmt"

	"github.com/Aritiaya50217/Poker/internal/domain"
	"github.com/Aritiaya50217/Poker/internal/service"
)

func Run() error {
	deck := domain.NewDeck()

	deck.Shuffle()

	players := []domain.Player{
		{
			Name: "Player 1",
			Hand: domain.NewHand(),
		},
		{
			Name: "Player 2",
			Hand: domain.NewHand(),
		},
		{
			Name: "Player 3",
			Hand: domain.NewHand(),
		},
		{
			Name: "Player 4",
			Hand: domain.NewHand(),
		},
	}

	for i := 0; i < 5; i++ {
		for j := range players {
			card, err := deck.Draw()
			if err != nil {
				return err
			}

			if err := players[j].Hand.Add(card); err != nil {
				return err
			}
		}
	}

	winner := 0
	best := service.Result{}

	for i, p := range players {
		result := service.Evaluate(p.Hand.Cards())
		fmt.Print(p.Name, ": ")

		for _, c := range p.Hand.Cards() {
			fmt.Print(c.String())
		}

		fmt.Printf(" -> %s\n", domain.RankName[result.Rank])

		if i == 0 {
			best = result

			continue
		}

		if service.Compare(result, best) == 1 {
			winner = i
			best = result
		}
	}

	fmt.Println()

	fmt.Println("Cards left in deck:", deck.Remaining())

	fmt.Printf("\n*** Winner is %s with %s! ***\n", players[winner].Name, domain.RankName[best.Rank])

	return nil
}
