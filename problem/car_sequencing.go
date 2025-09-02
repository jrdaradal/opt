package problem

import (
	"slices"
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/fn"
)

func CarSequencing(name string) *discrete.Problem {
	cfg := newCarSequencing(name)
	if cfg == nil {
		return nil
	}
	numCars, numOptions := len(cfg.cars), len(cfg.options)

	p := discrete.NewProblem(name)
	p.Goal = discrete.SATISFY

	p.Variables = discrete.Variables(cfg.cars)
	domain := discrete.IndexDomain(numCars)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllDiff constraint
	p.AddGlobalConstraint(allDiffConstraint)

	test := func(solution *discrete.Solution) bool {
		sequence := solution.AsSequence()
		optionSequence := make([][]bool, numOptions)
		for i := range numOptions {
			optionSequence[i] = make([]bool, numCars)
		}
		for seqIdx, x := range sequence {
			car := cfg.cars[x]
			for optionIdx, flag := range cfg.carOptions[car] {
				optionSequence[optionIdx][seqIdx] = flag
			}
		}
		for optionIdx, optionCfg := range cfg.options {
			maxCount, windowSize := optionCfg[0], optionCfg[1]
			limit := (numCars / windowSize) * windowSize
			for i := 0; i < limit; i += windowSize {
				window := optionSequence[optionIdx][i : i+windowSize]
				if fn.CountValue(window, true) > maxCount {
					return false
				}
			}
			window := optionSequence[optionIdx][limit:]
			if fn.CountValue(window, true) > maxCount {
				return false
			}
		}
		return true
	}
	p.AddGlobalConstraint(test)

	p.SolutionCore = discrete.MirroredSequence(cfg.cars)
	p.SolutionDisplay = discrete.DisplaySequence(cfg.cars)

	return p
}

type carSequenceCfg struct {
	options    [][2]int
	cars       []string
	carOptions map[string][]bool
}

func newCarSequencing(name string) *carSequenceCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 3 {
		return nil
	}
	cfg := &carSequenceCfg{
		options:    make([][2]int, 0),
		cars:       make([]string, 0),
		carOptions: make(map[string][]bool),
	}
	p := strings.Fields(lines[0])
	numOptions, numCars := fn.ParseInt(p[0]), fn.ParseInt(p[1])
	idx := 1
	optionIdx := make(map[string]int)
	for i := range numOptions {
		p = strings.Fields(lines[idx])
		name, maxCount, windowSize := p[0], fn.ParseInt(p[1]), fn.ParseInt(p[2])
		optionIdx[name] = i
		cfg.options = append(cfg.options, [2]int{maxCount, windowSize})
		idx++
	}
	for range numCars {
		p := strings.Fields(lines[idx])
		car, count := p[0], fn.ParseInt(p[1])
		cfg.cars = append(cfg.cars, slices.Repeat([]string{car}, count)...)
		flags := make([]bool, numOptions)
		for _, name := range p[2:] {
			i := optionIdx[name]
			flags[i] = true
		}
		cfg.carOptions[car] = flags
		idx++
	}
	return cfg
}
