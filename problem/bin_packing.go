package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func BinPacking(name string) *discrete.Problem {
	cfg := newBinPacking(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.weight)
	domain := discrete.RangeDomain(1, cfg.bins)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		sums := solution.PartitionSums(domain, cfg.weight)
		return fn.All(sums, func(sum int) bool {
			return sum <= cfg.capacity
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.UniqueValues
	p.SolutionCore = discrete.SortedPartition(domain, cfg.weight)
	p.SolutionDisplay = discrete.DisplayPartitions(domain, cfg.weight)

	return p
}

type binPackingCfg struct {
	capacity int
	bins     int
	weight   []int
}

func newBinPacking(name string) *binPackingCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &binPackingCfg{
		capacity: fn.ParseInt(lines[0]),
		bins:     fn.ParseInt(lines[1]),
		weight:   fn.Map(strings.Fields(lines[2]), fn.ParseInt),
	}
}
