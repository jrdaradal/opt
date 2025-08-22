package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func ResourceOptimization(name string) *discrete.Problem {
	cfg := newResourceOptimization(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MAXIMIZE

	p.Variables = discrete.Variables(cfg.resources)
	for i, variable := range p.Variables {
		p.Domain[variable] = discrete.RangeDomain(0, cfg.count[i])
	}

	test := func(solution *discrete.Solution) bool {
		count, cost := solution.Map, cfg.cost
		costs := fn.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cost[x]
		})
		return fn.Sum(costs) <= cfg.limit
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SumWeightedValues(p.Variables, cfg.value)
	p.SolutionDisplay = discrete.DisplayValues[int](p, nil)

	return p
}

type resourceOptCfg struct {
	limit     float64
	resources []string
	count     []int
	cost      []float64
	value     []float64
}

func newResourceOptimization(name string) *resourceOptCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 5 {
		return nil
	}
	return &resourceOptCfg{
		limit:     fn.ParseFloat(lines[0]),
		resources: strings.Fields(lines[1]),
		count:     fn.Map(strings.Fields(lines[2]), fn.ParseInt),
		cost:      fn.Map(strings.Fields(lines[3]), fn.ParseFloat),
		value:     fn.Map(strings.Fields(lines[4]), fn.ParseFloat),
	}
}
