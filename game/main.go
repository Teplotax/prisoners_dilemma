package main

import (
	"fmt"
	"github.com/Teplotax/prisoners_dilemma/strategy"
	strategyImpl "github.com/Teplotax/prisoners_dilemma/strategy/impl"
	"sort"
)

func main() {
	//Players
	player1 := strategyImpl.NewTitForTat()
	player2 := strategyImpl.NewRandom()
	player3 := strategyImpl.NewNasty()
	player4 := strategyImpl.NewTitForTwoTats()
	player5 := strategyImpl.NewFriedman()
	player6 := strategyImpl.NewJoss()
	player7 := strategyImpl.NewGraaskamp()

	players := []strategy.Strategy{
		&player1,
		&player2,
		&player3,
		&player4,
		&player5,
		&player6,
		&player7,
	}

	startGame(players, 5, 200)
}

func startGame(players []strategy.Strategy, games int32, rounds int32) {

	for game := int32(0); game < games; game++ {

		for i := 0; i < len(players)-1; i++ {
			for j := i + 1; j < len(players); j++ {
				resetScores(players)
				p1Choices := []bool{}
				p2Choices := []bool{}

				for n := int32(0); n < rounds; n++ {
					p1Choice := players[i].Play(p2Choices)
					p2Choice := players[j].Play(p1Choices)

					p1Choices = append(p1Choices, p1Choice)
					p2Choices = append(p2Choices, p2Choice)

					if p1Choice && p2Choice {
						players[i].AddScore(3)
						players[j].AddScore(3)
					} else if p1Choice && !p2Choice {
						players[i].AddScore(0)
						players[j].AddScore(5)
					} else if !p1Choice && p2Choice {
						players[i].AddScore(5)
						players[j].AddScore(0)
					} else if !p1Choice && !p2Choice {
						players[i].AddScore(1)
						players[j].AddScore(1)
					}

				}
			}
		}
	}

	printResults(players)
}

func resetScores(players []strategy.Strategy) {
	for _, player := range players {
		player.ResetCurrentScore()
	}
}

func printResults(players []strategy.Strategy) {
	results := make(map[string]int64)

	for i := 0; i < len(players); i++ {
		results[players[i].GetName()] = players[i].GetAverageScore()
	}

	var keys []string
	for key := range results {
		keys = append(keys, key)
	}

	// Step 4: Sort the keys based on their corresponding values in descending order
	sort.Slice(keys, func(i, j int) bool {
		return results[keys[i]] > results[keys[j]]
	})

	// Step 5: Print the sorted key-value pairs
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, results[key])
	}
}
