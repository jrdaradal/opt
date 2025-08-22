package problem

import (
	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func SetCover(name string) *discrete.Problem {
	subsets := newSetCover(name)
	if subsets == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(subsets.Names)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		covered := fn.BooleanMap(subsets.Universal, false)
		for _, x := range solution.AsSubset() {
			for _, item := range subsets.Subsets[x] {
				covered[item] = true
			}
		}
		return fn.AllTrue(fn.MapValues(covered))
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SubsetCount
	p.SolutionDisplay = discrete.DisplaySubset(subsets.Names)

	return p
}

func newSetCover(name string) *ds.Subsets {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	return ds.NewSubsets(lines[0], lines[1:])
}
