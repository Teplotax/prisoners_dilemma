package strategyImpl

type Nice struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewNice() Nice {
	return Nice{name: "Nice"}
}

func (s *Nice) Play(arr []bool) bool {
	return true
}

func (s *Nice) GetName() string {
	return s.name
}

func (s *Nice) GetScore() int64 {
	return s.score
}

func (s *Nice) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Nice) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Nice) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Nice) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
