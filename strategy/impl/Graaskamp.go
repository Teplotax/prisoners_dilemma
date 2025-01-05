package strategyImpl

type Graaskamp struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewGraaskamp() Graaskamp {
	return Graaskamp{name: "Graaskamp"}
}

func (s *Graaskamp) Play(arr []bool) bool {
	var choice bool
	if len(arr) == 0 {
		choice = true
	} else {

		if s.games%50 == 0 {
			choice = false
		} else {
			choice = arr[len(arr)-1]
		}
	}
	return choice
}

func (s *Graaskamp) GetName() string {
	return s.name
}

func (s *Graaskamp) GetScore() int64 {
	return s.score
}

func (s *Graaskamp) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *Graaskamp) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *Graaskamp) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *Graaskamp) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
