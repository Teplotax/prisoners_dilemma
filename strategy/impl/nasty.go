package strategyImpl

type Nasty struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewNasty() Nasty {
	return Nasty{name: "Nasty"}
}

func (s *Nasty) Play(arr []bool) bool {
	return false
}

func (s *Nasty) GetName() string {
	return s.name
}

func (s *Nasty) GetScore() int64 {
	return s.score
}

func (s *Nasty) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Nasty) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Nasty) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Nasty) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
