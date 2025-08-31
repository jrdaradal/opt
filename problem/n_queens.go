package problem

import (
	"fmt"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
)

func NQueens(n int) *discrete.Problem {
	name := fmt.Sprintf("nqueens%d", n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.SATISFY

	p.Variables = discrete.RangeVariables(1, n)
	domain := discrete.RangeDomain(1, n)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// NoRowConflict
	p.AddGlobalConstraint(allDiffConstraint)

	// NoDiagonalConflict
	test2 := func(solution *discrete.Solution) bool {
		row := solution.Map
		occupied := ds.NewSet[ds.Coords]()
		for _, x := range p.Variables {
			occupied.Add(ds.Coords{row[x], x})
		}
		for _, x := range p.Variables {
			c := ds.Coords{row[x], x}
			if hasDiagonalConflict(c, occupied, n) {
				return false
			}
		}
		return true
	}
	p.AddGlobalConstraint(test2)

	p.SolutionDisplay = discrete.DisplayValues[int](p, nil)

	return p
}

func hasDiagonalConflict(c ds.Coords, occupied *ds.Set[ds.Coords], n int) bool {
	row, col := c.Tuple()
	// Upper Left
	for y, x := row-1, col-1; y >= 1 && x >= 1; {
		if occupied.Contains(ds.Coords{y, x}) {
			return true
		}
		x--
		y--
	}
	// Upper Right
	for y, x := row-1, col+1; y >= 1 && x <= n; {
		if occupied.Contains(ds.Coords{y, x}) {
			return true
		}
		x++
		y--
	}
	// Bottom Left
	for y, x := row+1, col-1; y <= n && x >= 1; {
		if occupied.Contains(ds.Coords{y, x}) {
			return true
		}
		x--
		y++
	}
	// Bottom Right
	for y, x := row+1, col+1; y <= n && x <= n; {
		if occupied.Contains(ds.Coords{y, x}) {
			return true
		}
		x++
		y++
	}

	return false
}
