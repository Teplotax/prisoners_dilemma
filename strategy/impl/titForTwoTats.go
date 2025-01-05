package strategyImpl

type TitForTwoTats struct {
	name         string
	score        int64
	currentScore int64
	games        int64
}

func NewTitForTwoTats() TitForTwoTats {
	return TitForTwoTats{name: "Tit for Two Tats"}
}

func (s *TitForTwoTats) Play(arr []bool) bool {
	var choice bool
	if len(arr) <= 1 {
		choice = true
	} else {
		if !arr[len(arr)-1] && !arr[len(arr)-2] {
			choice = false
		} else {
			choice = true
		}
	}

	return choice
}

func (s *TitForTwoTats) GetName() string {
	return s.name
}

func (s *TitForTwoTats) GetScore() int64 {
	return s.score
}

func (s *TitForTwoTats) GetCurrentScore() int64 {
	return s.currentScore
}

func (s *TitForTwoTats) GetAverageScore() int64 {
	return s.score / s.games
}

func (s *TitForTwoTats) AddScore(points int64) {
	s.currentScore += points
	s.score += points
}

func (s *TitForTwoTats) ResetCurrentScore() {
	s.currentScore = 0
	s.games++
}
