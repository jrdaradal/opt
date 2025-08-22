package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func EdgeColoring(name string) *discrete.Problem {
	graph, colors := newEdgeColoring(name)
	if graph == nil || colors == nil {
		return nil
	}

	edgeNames := fn.Map(graph.Edges, fn.ToString)
	indexOf := make(map[string]int)
	for i, edgeName := range edgeNames {
		indexOf[edgeName] = i
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.MapDomain(colors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		color := solution.Map
		return fn.All(graph.Vertices, func(vertex ds.Vertex) bool {
			edgeColors := fn.Map(graph.ConnectedEdges(vertex), func(edge ds.Edge) int {
				x := indexOf[edge.String()]
				return color[x]
			})
			return ds.AllUnique(edgeColors)
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = discrete.UniqueValues
	p.SolutionCore = discrete.LookupValueOrder(p)
	p.SolutionDisplay = discrete.DisplayValues(p, colors)

	return p
}

func newEdgeColoring(name string) (*ds.Graph, []string) {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 3 {
		return nil, nil
	}
	graph := ds.NewGraph(lines[0], lines[1])
	colors := strings.Fields(lines[2])
	return graph, colors
}
