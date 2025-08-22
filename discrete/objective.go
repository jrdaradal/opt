package discrete

import (
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

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

func UniqueValues(solution *Solution) Score {
	uniqueValues := ds.SetFrom(solution.Values())
	solution.Score = Score(uniqueValues.Len())
	return solution.Score
}

func SumWeightedValues(variables []Variable, value []float64) ObjectiveFunc {
	return func(solution *Solution) Score {
		count := solution.Map
		solution.Score = fn.Sum(fn.Map(variables, func(x Variable) Score {
			return float64(count[x]) * value[x]
		}))
		return solution.Score
	}
}
