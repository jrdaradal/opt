package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func SubsetSum(name string) *discrete.Problem {
	target, numbers := newSubsetSum(name)
	if target == 0 || numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(numbers)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		total := fn.SumValues(solution.AsSubset(), numbers)
		if p.IsSatisfactionProblem() {
			return total == target
		} else {
			return total <= target
		}
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		total := fn.SumValues(solution.AsSubset(), numbers)
		solution.Score = discrete.Score(target - total)
		return solution.Score
	}

	p.SolutionDisplay = discrete.DisplaySubset(numbers)

	return p
}

func newSubsetSum(name string) (int, []int) {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 2 {
		return 0, nil
	}
	target := fn.ParseInt(lines[0])
	numbers := fn.Map(strings.Fields(lines[1]), fn.ParseInt)
	return target, numbers
}
