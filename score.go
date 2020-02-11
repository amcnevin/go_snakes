package main

//Score container
type Score struct {
	score int
}

func initScore() Score {
	score := Score{}
	score.score = 0
	return score
}

func incrScore(score *Score) {
	score.score++
}
