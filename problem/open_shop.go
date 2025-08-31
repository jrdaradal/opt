package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func OpenShop(name string) *discrete.Problem {
	cfg := newOpenShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.tasks)
	domain := discrete.RangeDomain(0, cfg.maxMakespan)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// 1) Job tasks no overlap
	p.AddGlobalConstraint(noJobTaskOverlap(cfg))

	// 2) No machine overlap
	p.AddGlobalConstraint(noMachineOverlap(cfg))

	p.ObjectiveFunc = scheduleMakespan(cfg.tasks)
	p.SolutionDisplay = discrete.DisplayShopSchedule(cfg.tasks, cfg.machines)

	return p
}

func newOpenShop(name string) *shopSchedCfg {
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
	for jobID, line := range lines[1:] {
		job := &ds.Job{
			ID:    jobID,
			Tasks: make([]*ds.Task, 0),
		}
		for taskID, d := range strings.Fields(line) {
			text := ds.TaskString(cfg.machines[taskID], d)
			task := ds.NewTask(text, jobID, taskID)
			job.Tasks = append(job.Tasks, task)
		}
		cfg.jobs = append(cfg.jobs, job)
		cfg.tasks = append(cfg.tasks, job.Tasks...)
		totalDuration += job.TotalDuration()
	}
	cfg.maxMakespan = totalDuration
	return cfg
}
