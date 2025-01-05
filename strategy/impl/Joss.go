package strategyImpl

import (
	"math/rand"
	"time"
)

type Joss struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewJoss() Joss {
	return Joss{name: "Joss"}
}

func (s *Joss) Play(arr []bool) bool {
	var choice bool
	if len(arr) == 0 {
		choice = true
	} else {
		rand.Seed(time.Now().UnixNano())

		// Generate a random number between 0 and 100
		randomNum := rand.Intn(100)

		// Return false if the number is <= 10 (10% chance)
		if randomNum <= 10 {
			choice = false
		} else {
			choice = arr[len(arr)-1]
		}
	}
	return choice
}

func (s *Joss) GetName() string {
	return s.name
}

func (s *Joss) GetScore() int64 {
	return s.score
}

func (s *Joss) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Joss) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Joss) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Joss) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
