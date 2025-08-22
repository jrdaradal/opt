package problem

import (
	"fmt"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func MagicSeries(n int) *discrete.Problem {
	name := fmt.Sprintf("magicseries%d", n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.SATISFY

	p.Variables = discrete.RangeVariables(0, n)
	domain := discrete.RangeDomain(0, n)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		value := solution.Map
		count := fn.TallyValues(solution.Map, domain)
		return fn.All(p.Variables, func(x discrete.Variable) bool {
			return value[x] == count[x]
		})
	}
	p.AddGlobalConstraint(test)

	p.SolutionDisplay = discrete.DisplayValues[int](p, nil)

	return p
}
