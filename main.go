package main

import (
	"log"
	"math/rand"
	"time"
)

const workoutLength = 20

func main() {
	rand.Seed(time.Now().Unix())

	cardSet := initSet()
	var randomCard string
	var randIndex int
	for i := 0; i < workoutLength; i++ {
		randIndex = rand.Intn(len(cardSet.Cards))
		randomCard = cardSet.Cards[randIndex]
		log.Printf("%d. %s", i+1, cardSet.getCardTranslation(randomCard))
	}
}
