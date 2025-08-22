package problem

import (
	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func Clique(name string) *discrete.Problem {
	graph := newClique(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MAXIMIZE

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		vertices := fn.Map(solution.AsSubset(), func(x int) string {
			return graph.Vertices[x]
		})
		return graph.IsClique(vertices)
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.SubsetCount
	p.SolutionDisplay = discrete.DisplaySubset(graph.Vertices)

	return p
}

func newClique(name string) *ds.Graph {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return ds.NewGraph(lines[0], lines[1])
}
