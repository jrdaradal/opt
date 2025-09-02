package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func SceneAllocation(name string) *discrete.Problem {
	cfg := newSceneAllocation(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.scenes)
	domain := discrete.MapDomain(cfg.days)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		return fn.AllIndex(solution.AsPartitions(domain), func(day int, scenes []discrete.Variable) bool {
			limits := cfg.days[day]
			minScene, maxScene := limits[0], limits[1]
			currScene := len(scenes)
			return minScene <= currScene && currScene <= maxScene
		})
	}
	p.AddGlobalConstraint(test)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		var score discrete.Score = 0
		for _, scenes := range solution.AsPartitions(domain) {
			dailyActors := ds.NewSet[string]()
			for _, x := range scenes {
				scene := cfg.scenes[x]
				dailyActors.AddItems(cfg.sceneActors[scene])
			}
			score += fn.SumMapValues(dailyActors.Items(), cfg.dailyCost)
		}
		solution.Score = score
		return solution.Score
	}

	p.SolutionCore = discrete.SortedPartition(domain, cfg.scenes)
	p.SolutionDisplay = discrete.DisplayPartitions(domain, cfg.scenes)

	return p
}

type sceneCfg struct {
	days        [][2]int
	dailyCost   map[string]float64
	scenes      []string
	sceneActors map[string][]string
}

func newSceneAllocation(name string) *sceneCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 5 {
		return nil
	}
	cfg := &sceneCfg{
		days:        make([][2]int, 0),
		dailyCost:   make(map[string]float64),
		scenes:      make([]string, 0),
		sceneActors: make(map[string][]string),
	}
	p := strings.Fields(lines[0])
	numDays, numActors, numScenes := fn.ParseInt(p[0]), fn.ParseInt(p[1]), fn.ParseInt(p[2])
	minScenes := fn.Map(strings.Fields(lines[1]), fn.ParseInt)
	maxScenes := fn.Map(strings.Fields(lines[2]), fn.ParseInt)
	for i := range numDays {
		cfg.days = append(cfg.days, [2]int{minScenes[i], maxScenes[i]})
	}
	idx := 3
	for range numActors {
		p = strings.Fields(lines[idx])
		name, cost := p[0], fn.ParseFloat(p[1])
		cfg.dailyCost[name] = cost
		idx++
	}
	for range numScenes {
		p = strings.Fields(lines[idx])
		scene, actors := p[0], p[1:]
		cfg.scenes = append(cfg.scenes, scene)
		cfg.sceneActors[scene] = actors
		idx++
	}
	return cfg
}
