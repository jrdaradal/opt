package problem

import (
	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func ExactCover(name string) *discrete.Problem {
	subsets := newExactCover(name)
	if subsets == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.SATISFY

	p.Variables = discrete.Variables(subsets.Names)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		count := fn.NewCounter(subsets.Universal)
		for _, x := range solution.AsSubset() {
			for _, item := range subsets.Subsets[x] {
				count[item] += 1
			}
		}
		return fn.AllEqual(fn.MapValues(count), 1)
	}
	p.AddGlobalConstraint(test)

	p.SolutionDisplay = discrete.DisplaySubset(subsets.Names)

	return p
}

func newExactCover(name string) *ds.Subsets {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	return ds.NewSubsets(lines[0], lines[1:])
}
