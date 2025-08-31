package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func CarPainting(name string) *discrete.Problem {
	cfg := newCarPainting(name)
	if cfg == nil {
		return nil
	}
	numCars := len(cfg.cars)

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.cars)
	minLimit, maxLimit := 0, numCars-1
	for _, variable := range p.Variables {
		first := max(minLimit, variable-cfg.maxShift)
		last := min(maxLimit, variable+cfg.maxShift)
		p.Domain[variable] = discrete.RangeDomain(first, last)
	}

	// AllDiff constraint
	p.AddGlobalConstraint(allDiffConstraint)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		colorSequence := make([]string, numCars)
		for i, x := range solution.AsSequence() {
			colorSequence[i] = cfg.cars[x]
		}
		changes := countColorChanges(colorSequence)
		solution.Score = discrete.Score(changes)
		return solution.Score
	}

	p.SolutionCore = func(solution *discrete.Solution) string {
		sequence := make([]ds.StrInt, numCars)
		for i, x := range solution.AsSequence() {
			sequence[i] = ds.StrInt{Str: cfg.cars[x], Int: x}
		}
		var prevColor string
		groups := make([][]int, 0)
		group := make([]int, 0)
		for i, item := range sequence {
			currColor := item.Str
			if i > 0 && prevColor != currColor {
				groups = append(groups, group)
				group = []int{item.Int}
			} else {
				group = append(group, item.Int)
			}
			prevColor = currColor
		}
		groups = append(groups, group)
		output := fn.Map(groups, func(group []int) string {
			slices.Sort(group)
			return strings.Join(fn.Map(group, fn.IntToString), " ")
		})
		return strings.Join(output, "|")
	}

	p.SolutionDisplay = func(solution *discrete.Solution) string {
		sequence := make([]ds.StrInt, numCars)
		for i, x := range solution.AsSequence() {
			sequence[i] = ds.StrInt{Str: cfg.cars[x], Int: x}
		}
		output := make([]string, 0)
		var prevColor string
		for i, item := range sequence {
			currColor := item.Str
			if i > 0 && prevColor != currColor {
				output = append(output, "|")
			}
			output = append(output, fmt.Sprintf("%d:%s", item.Int, item.Str))
			prevColor = currColor
		}
		return strings.Join(output, " ")
	}

	return p
}

type carPaintCfg struct {
	maxShift int
	cars     []string
}

func newCarPainting(name string) *carPaintCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return &carPaintCfg{
		maxShift: fn.ParseInt(lines[0]),
		cars:     strings.Fields(lines[1]),
	}
}
