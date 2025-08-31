package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func BinaryPaintShop(name string) *discrete.Problem {
	cfg := newBinaryPaintShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.RangeVariables(0, cfg.numCars-1)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Set first car to 0
	car0 := p.Variables[0]
	p.Domain[car0] = []discrete.Value{0}

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		// Initialize current colors of cars from solution
		color := make([]int, cfg.numCars)
		for x, c := range solution.Map {
			color[x] = c
		}
		colorSequence := make([]int, len(cfg.sequence))
		for i, x := range cfg.sequence {
			colorSequence[i] = color[x]
			color[x] = (color[x] + 1) % 2 // flip
		}
		prevColor, changes := 0, 0
		for i, currColor := range colorSequence {
			if i > 0 && prevColor != currColor {
				changes += 1
			}
			prevColor = currColor
		}
		solution.Score = discrete.Score(changes)
		return solution.Score
	}

	p.SolutionDisplay = discrete.DisplayValues[int](p, nil)

	return p
}

type binaryPaintCfg struct {
	numCars  int
	sequence []int
}

func newBinaryPaintShop(name string) *binaryPaintCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return &binaryPaintCfg{
		numCars:  fn.ParseInt(lines[0]),
		sequence: fn.Map(strings.Fields(lines[1]), fn.ParseInt),
	}
}
