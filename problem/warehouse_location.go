package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func WarehouseLocation(name string) *discrete.Problem {
	cfg := newWarehouseLocation(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.stores)
	domain := discrete.MapDomain(cfg.warehouses)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		usage := fn.TallyValues(solution.Map, domain)
		return fn.AllTrue(fn.MapIndex(domain, func(i, w int) bool {
			return usage[w] <= cfg.capacity[i]
		}))
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		var totalCost discrete.Score = 0
		// Fixed cost
		usage := fn.TallyValues(solution.Map, domain)
		for i, w := range domain {
			if usage[w] > 0 {
				totalCost += cfg.fixedCost[i]
			}
		}
		// Supply cost
		for x, w := range solution.Map {
			totalCost += cfg.supplyCost[x][w]
		}
		solution.Score = totalCost
		return solution.Score
	}

	p.SolutionDisplay = discrete.DisplayPartitions(domain, cfg.stores)

	return p
}

type warehouseCfg struct {
	warehouses []string // size M
	stores     []string // size N
	capacity   []int
	fixedCost  []float64
	supplyCost [][]float64 // 1 row per store => vector of size M (cost per warehouse)
}

func newWarehouseLocation(name string) *warehouseCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 4 {
		return nil
	}
	cfg := &warehouseCfg{
		warehouses: strings.Fields(lines[0]),
		capacity:   fn.Map(strings.Fields(lines[1]), fn.ParseInt),
		fixedCost:  fn.Map(strings.Fields(lines[2]), fn.ParseFloat),
		stores:     make([]string, 0),
		supplyCost: make([][]float64, 0),
	}
	for _, line := range lines[3:] {
		parts := fn.CleanSplit(line, ":")
		cost := fn.Map(strings.Fields(parts[1]), fn.ParseFloat)
		cfg.stores = append(cfg.stores, parts[0])
		cfg.supplyCost = append(cfg.supplyCost, cost)
	}
	return cfg
}
