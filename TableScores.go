package main

type TableScores struct {
	name  string
	score int
}

func createNewTableScores(name string, score int) TableScores {
	newTableScores := TableScores{
		name:  name,
		score: score,
	}

	return newTableScores
}
