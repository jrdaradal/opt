package problem

import (
	"strings"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
	"github.com/jrdaradal/opt/internal/fn"
)

func FlowShop(name string) *discrete.Problem {
	cfg := newFlowShop(name)
	if cfg == nil {
		return nil
	}
	numJobs := len(cfg.jobs)
	numMachines := len(cfg.machines)

	p := discrete.NewProblem(name)
	p.Goal = discrete.MINIMIZE

	p.Variables = discrete.Variables(cfg.jobs)
	domain := discrete.IndexDomain(numJobs)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllDiff constraint
	p.AddGlobalConstraint(allDiffConstraint)

	p.ObjectiveFunc = func(solution *discrete.Solution) discrete.Score {
		sequence := solution.AsSequence()
		end := make(map[ds.Coords]int)
		for m := range numMachines {
			for i, x := range sequence {
				task := cfg.jobs[x].Tasks[m]
				above := end[ds.Coords{m - 1, i}]
				prev := end[ds.Coords{m, i - 1}]
				// Pick the later ending: above or prev
				// above = previous task on the same job (can only process one job task at a time)
				// prev = previous task on the same machine (machine can only process one task at a time)
				start := max(above, prev)
				end[ds.Coords{m, i}] = start + task.Duration
			}
		}
		y, x := numMachines-1, numJobs-1
		solution.Score = discrete.Score(end[ds.Coords{y, x}])
		return solution.Score
	}

	jobIDs := fn.Map(cfg.jobs, func(job *ds.Job) int {
		return job.ID
	})
	p.SolutionDisplay = discrete.DisplaySequence(jobIDs)

	return p
}

func newFlowShop(name string) *shopSchedCfg {
	lines, err := fn.ProblemData(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &shopSchedCfg{
		machines: strings.Fields(lines[0]),
		jobs:     make([]*ds.Job, 0),
		tasks:    make([]*ds.Task, 0),
	}
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
	}
	return cfg
}
