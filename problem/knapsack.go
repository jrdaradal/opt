package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func Knapsack(name string) *discrete.Problem {
	cfg := newKnapsack(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MAXIMIZE

	p.Variables = discrete.Variables(cfg.items)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		count, weight := solution.Map, cfg.weight
		weights := fn.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * weight[x]
		})
		return fn.Sum(weights) <= cfg.capacity
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SumWeightedValues(p.Variables, cfg.value)
	p.SolutionDisplay = discrete.DisplaySubset(cfg.items)

	return p
}

type knapsackCfg struct {
	capacity float64
	items    []string
	weight   []float64
	value    []float64
}

func newKnapsack(name string) *knapsackCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 4 {
		return nil
	}
	return &knapsackCfg{
		capacity: fn.ParseFloat(lines[0]),
		items:    strings.Fields(lines[1]),
		weight:   fn.Map(strings.Fields(lines[2]), fn.ParseFloat),
		value:    fn.Map(strings.Fields(lines[3]), fn.ParseFloat),
	}
}
