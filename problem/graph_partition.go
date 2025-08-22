package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func GraphPartition(name string) *discrete.Problem {
	cfg, graph := newGraphPartition(name)
	if cfg == nil || graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.RangeDomain(1, cfg.numPartitions)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		partitionSizes := fn.TallyValues(solution.Map, domain)
		return fn.All(fn.MapValues(partitionSizes), func(size int) bool {
			return size >= cfg.minPartitionSize
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		var score discrete.Score = 0
		group := solution.Map
		for i, edge := range graph.Edges {
			x1, x2 := graph.IndexOf(edge[0]), graph.IndexOf(edge[1])
			if group[x1] != group[x2] {
				score += cfg.edgeWeight[i]
			}
		}
		solution.Score = score
		return solution.Score
	}

	p.SolutionCore = discrete.SortedPartition(domain, graph.Vertices)
	p.SolutionDisplay = discrete.DisplayPartitions(domain, graph.Vertices)

	return p
}

type graphPartitionCfg struct {
	numPartitions    int
	minPartitionSize int
	edgeWeight       []float64
}

func newGraphPartition(name string) (*graphPartitionCfg, *ds.Graph) {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 5 {
		return nil, nil
	}
	graph := ds.NewGraph(lines[2], lines[3])
	cfg := &graphPartitionCfg{
		numPartitions:    fn.ParseInt(lines[0]),
		minPartitionSize: fn.ParseInt(lines[1]),
		edgeWeight:       fn.Map(strings.Fields(lines[4]), fn.ParseFloat),
	}
	return cfg, graph
}
