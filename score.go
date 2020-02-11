package main

//Score container
type Score struct {
	score int
}

// Initiatize Score
func initScore() Score {
	score := Score{}
	score.score = 0
	return score
}

// Increment Score
func incrScore(score *Score) {
	score.score++
}
