package strategyImpl

type TitForTat struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewTitForTat() TitForTat {
	return TitForTat{name: "Tit for Tat"}
}

func (s *TitForTat) Play(arr []bool) bool {
	var choice bool
	if len(arr) == 0 {
		choice = true
	} else {
		choice = arr[len(arr)-1]
	}

	return choice
}

func (s *TitForTat) GetName() string {
	return s.name
}

func (s *TitForTat) GetScore() int64 {
	return s.score
}

func (s *TitForTat) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *TitForTat) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *TitForTat) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *TitForTat) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
