package strategyImpl

type Friedman struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewFriedman() Friedman {
	return Friedman{name: "Friedman"}
}

func (s *Friedman) Play(arr []bool) bool {
	for _, val := range arr {
		if !val {
			return false // Return false if any value is false
		}
	}
	return true
}

func (s *Friedman) GetName() string {
	return s.name
}

func (s *Friedman) GetScore() int64 {
	return s.score
}

func (s *Friedman) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Friedman) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Friedman) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Friedman) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
