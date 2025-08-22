package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func GraphColoring(name string) *discrete.Problem {
	graph, colors, asMap := newGraphColoring(name)
	if graph == nil || colors == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.MapDomain(colors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		color := solution.Map
		return fn.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf(edge[0]), graph.IndexOf(edge[1])
			return color[x1] != color[x2]
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.UniqueValues
	p.SolutionCore = discrete.LookupValueOrder(p)
	if asMap {
		p.SolutionDisplay = discrete.DisplayMap(p, graph.Vertices, colors)
	} else {
		p.SolutionDisplay = discrete.DisplayValues(p, colors)
	}

	return p
}

func newGraphColoring(name string) (*ds.Graph, []string, bool) {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 3 {
		return nil, nil, false
	}
	graph := ds.NewGraph(lines[0], lines[1])
	colors := strings.Fields(lines[2])
	asMap := strings.HasSuffix(name, "m")
	return graph, colors, asMap
}
