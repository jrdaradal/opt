package problem

import (
	"sort"
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func LIS(name string) *discrete.Problem {
	sequence := newLIS(name)
	if sequence == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MAXIMIZE

	p.Variables = discrete.Variables(sequence)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		subset := solution.AsSubset()
		sort.Ints(subset)
		subseq := fn.Map(subset, func(x discrete.Variable) int {
			return sequence[x]
		})
		for i := 1; i < len(subseq); i++ {
			prev, curr := subseq[i-1], subseq[i]
			if prev >= curr {
				return false
			}
		}
		return true
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SubsetCount
	p.SolutionDisplay = discrete.DisplaySubset(sequence)

	return p
}

func newLIS(name string) []int {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return fn.Map(strings.Fields(lines[0]), fn.ParseInt)
}
