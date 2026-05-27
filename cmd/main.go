package main

import (
	"log"

	"github.com/Aritiaya50217/Poker/internal/usecase"
)

func main() {
	if err := usecase.Run(); err != nil {
		log.Fatal(err)
	}
}
