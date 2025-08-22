package problem

import (
	"fmt"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func LangfordPair(n int) *discrete.Problem {
	name := fmt.Sprintf("langford%d", n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.SATISFY

	numPositions := n * 2
	numbers := make([]int, 0, numPositions)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i, i)
	}

	p.Variables = discrete.Variables(numbers)
	domain := discrete.IndexDomain(numPositions)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllDiff constraint
	test1 := func(solution *discrete.Solution) bool {
		return ds.AllUnique(solution.Values())
	}
	p.AddGlobalConstraint(test1)

	// Distance constraint
	test2 := func(solution *discrete.Solution) bool {
		index := solution.Map
		for x := 0; x < len(p.Variables); x += 2 {
			number := (x / 2) + 1
			gap := fn.Abs(index[x+1]-index[x]) - 1
			if gap != number {
				return false
			}
		}
		return true
	}
	p.AddGlobalConstraint(test2)

	p.SolutionCore = discrete.MirroredSequence(numbers)
	p.SolutionDisplay = discrete.DisplaySequence(numbers)

	return p
}
