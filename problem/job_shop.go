package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func JobShop(name string) *discrete.Problem {
	cfg := newJobShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	variableID := 0
	jobTasks := make(map[int][]discrete.Variable)
	for jobID, job := range cfg.jobs {
		jobTasks[jobID] = make([]discrete.Variable, 0)
		for taskID, task := range job.Tasks {
			first, after := job.TaskMargins(taskID)
			last := cfg.maxMakespan - after - task.Duration
			variable := discrete.Variable(variableID)
			p.Variables = append(p.Variables, variable)
			p.Domain[variable] = discrete.RangeDomain(first, last)
			jobTasks[jobID] = append(jobTasks[jobID], variable)
			variableID++
		}
	}

	// 1) Job tasks in order and no overlap
	test := func(solution *discrete.Solution) bool {
		for _, variables := range jobTasks {
			for i := range len(variables) - 1 {
				curr, next := variables[i], variables[i+1]
				start1, start2 := solution.Map[curr], solution.Map[next]
				end1 := start1 + cfg.tasks[curr].Duration
				// Not in order or has overlap
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	}
	p.AddGlobalConstraint(test)

	// 2) No machine overlap
	p.AddGlobalConstraint(noMachineOverlap(cfg))

	p.ObjectiveFunc = scheduleMakespan(cfg.tasks)
	p.SolutionDisplay = discrete.DisplayShopSchedule(cfg.tasks, cfg.machines)

	return p
}

type shopSchedCfg struct {
	machines    []string
	jobs        []*ds.Job
	tasks       []*ds.Task
	maxMakespan int
}

func newJobShop(name string) *shopSchedCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &shopSchedCfg{
		machines: strings.Fields(lines[0]),
		jobs:     make([]*ds.Job, 0),
		tasks:    make([]*ds.Task, 0),
	}
	totalDuration := 0
	for _, line := range lines[1:] {
		parts := fn.CleanSplit(line, "=")
		if len(parts) != 2 {
			continue
		}
		job := ds.NewJob(parts[1], parts[0])
		cfg.jobs = append(cfg.jobs, job)
		cfg.tasks = append(cfg.tasks, job.Tasks...)
		totalDuration += job.TotalDuration()
	}
	cfg.maxMakespan = totalDuration
	return cfg
}
