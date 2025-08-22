package discrete

type Score = float64
type Goal string

const (
	MAXIMIZE Goal = "max"
	MINIMIZE Goal = "min"
	SATISFY  Goal = "sat"
)

type ObjectiveFunc func(*Solution) Score

func SubsetCount(solution *Solution) Score {
	solution.Score = Score(solution.SubsetCount())
	return solution.Score
}
