package strategyImpl

import (
	"math/rand"
	"time"
)

type Random struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewRandom() Random {
	return Random{name: "Random"}
}

func (s *Random) Play(arr []bool) bool {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	choice := rand.Intn(2) == 1
	return choice
}

func (s *Random) GetName() string {
	return s.name
}

func (s *Random) GetScore() int64 {
	return s.score
}

func (s *Random) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Random) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Random) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Random) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
