package strategy

type Strategy interface {
	Play(arr []bool) bool
	GetName() string
	GetScore() int64
	GetCurrentScore() int64
	GetAverageScore() int64
	AddScore(points int64)
	ResetCurrentScore()
}
