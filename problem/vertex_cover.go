package problem

import (
	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func VertexCover(name string) *discrete.Problem {
	graph := newVertexCover(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		used := solution.Map
		return fn.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf(edge[0]), graph.IndexOf(edge[1])
			return used[x1]+used[x2] > 0 // at least one is covered
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SubsetCount
	p.SolutionDisplay = discrete.DisplaySubset(graph.Vertices)

	return p
}

func newVertexCover(name string) *ds.Graph {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return ds.NewGraph(lines[0], lines[1])
}
