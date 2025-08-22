package problem

import (
	"cmp"
	"slices"
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func ActivitySelection(name string) *discrete.Problem {
	cfg := newActivitySelection(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MAXIMIZE

	p.Variables = discrete.Variables(cfg.activities)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		selected := solution.AsSubset()
		numSelected := len(selected)
		if numSelected <= 1 {
			return true
		}
		start, end := cfg.start, cfg.end
		slices.SortFunc(selected, func(x1, x2 int) int {
			return cmp.Compare(start[x1], start[x2])
		})
		for i := 1; i < numSelected; i++ {
			prev, curr := selected[i-1], selected[i]
			if end[prev] > start[curr] {
				return false
			}
		}
		return true
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SubsetCount
	p.SolutionDisplay = discrete.DisplaySubset(cfg.activities)

	return p
}

type activitySelectionCfg struct {
	activities []string
	start      []float64
	end        []float64
}

func newActivitySelection(name string) *activitySelectionCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &activitySelectionCfg{
		activities: strings.Fields(lines[0]),
		start:      fn.Map(strings.Fields(lines[1]), fn.ParseFloat),
		end:        fn.Map(strings.Fields(lines[2]), fn.ParseFloat),
	}
}
