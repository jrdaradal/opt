package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func NumberPartition(name string) *discrete.Problem {
	numbers := newNumberPartition(name)
	if numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(numbers)
	domain := discrete.RangeDomain(1, 2)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		if p.IsOptimizationProblem() {
			return true // don't test if optimization problem
		}
		sums := solution.PartitionSums(domain, numbers)
		return ds.AllSame(sums)
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		sums := solution.PartitionSums(domain, numbers)
		solution.Score = discrete.Score(fn.Abs(sums[0] - sums[1]))
		return solution.Score
	}

	p.SolutionCore = discrete.SortedPartition(domain, numbers)
	p.SolutionDisplay = discrete.DisplayPartitions(domain, numbers)

	return p
}

func newNumberPartition(name string) []int {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return fn.Map(strings.Fields(lines[0]), fn.ParseInt)
}
