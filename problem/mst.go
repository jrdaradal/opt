package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func MST(name string) *discrete.Problem {
	graph, edgeWeight := newMST(name)
	if graph == nil || edgeWeight == nil {
		return nil
	}
	edgeNames := fn.Map(graph.Edges, fn.ToString)

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Constraint: all vertices are spanned
	test1 := func(solution *discrete.Solution) bool {
		spanned := fn.BooleanMap(graph.Vertices, false)
		for _, x := range solution.AsSubset() {
			v1, v2 := graph.Edges[x].Tuple()
			spanned[v1] = true
			spanned[v2] = true

		}
		return fn.AllTrue(fn.MapValues(spanned))
	}
	p.AddGlobalConstraint(test1)

	// Constraint: solution forms a tree
	test2 := func(solution *discrete.Solution) bool {
		edges := fn.Map(solution.AsSubset(), func(x discrete.Variable) ds.Edge {
			return graph.Edges[x]
		})
		if len(edges) == 0 {
			return false
		}
		activeEdges := ds.SetFrom(edges)
		startVertex := edges[0][0] // first edge's first vertex
		reachable := ds.SetFrom(graph.BFSTraversal(startVertex, activeEdges))
		vertexSet := ds.SetFrom(graph.Vertices)
		return vertexSet.Diff(reachable).IsEmpty()
	}
	p.AddGlobalConstraint(test2)

	p.ObjectiveFunc = discrete.SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionDisplay = discrete.DisplaySubset(edgeNames)

	return p
}

func newMST(name string) (*ds.Graph, []float64) {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 3 {
		return nil, nil
	}
	graph := ds.NewGraph(lines[0], lines[1])
	edgeWeight := fn.Map(strings.Fields(lines[2]), fn.ParseFloat)
	return graph, edgeWeight
}
