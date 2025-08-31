package problem

import (
	"slices"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/internal/ds"
)

func allDiffConstraint(solution *discrete.Solution) bool {
	return ds.AllUnique(solution.Values())
}

func noMachineOverlap(cfg *shopSchedCfg) discrete.ConstraintFunc {
	return func(solution *discrete.Solution) bool {
		machineSched := make(map[string][]ds.TimeRange)
		for _, machine := range cfg.machines {
			machineSched[machine] = make([]ds.TimeRange, 0)
		}
		for x, start := range solution.Map {
			task := cfg.tasks[x]
			sched := ds.TimeRange{start, start + task.Duration}
			machine := task.Machine
			machineSched[machine] = append(machineSched[machine], sched)
		}
		for _, scheds := range machineSched {
			slices.SortFunc(scheds, ds.SortByStartTime)
			for i := range len(scheds) - 1 {
				curr, next := scheds[i], scheds[i+1]
				start1, end1 := curr.Tuple()
				start2 := next[0]
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	}
}

func noJobTaskOverlap(cfg *shopSchedCfg) discrete.ConstraintFunc {
	return func(solution *discrete.Solution) bool {
		jobSched := make(map[string][]ds.TimeRange)
		for _, job := range cfg.jobs {
			jobSched[job.ID] = make([]ds.TimeRange, 0)
		}
		for x, start := range solution.Map {
			task := cfg.tasks[x]
			sched := ds.TimeRange{start, start + task.Duration}
			jobID := task.JobID
			jobSched[jobID] = append(jobSched[jobID], sched)
		}
		for _, scheds := range jobSched {
			slices.SortFunc(scheds, ds.SortByStartTime)
			for i := range len(scheds) - 1 {
				curr, next := scheds[i], scheds[i+1]
				start1, end1 := curr.Tuple()
				start2 := next[0]
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	}
}

func scheduleMakespan(tasks []*ds.Task) discrete.ObjectiveFunc {
	return func(solution *discrete.Solution) discrete.Score {
		makespan := 0
		for x, start := range solution.Map {
			task := tasks[x]
			end := start + task.Duration
			makespan = max(makespan, end)
		}
		solution.Score = discrete.Score(makespan)
		return solution.Score
	}
}

func countColorChanges[T comparable](colorSequence []T) int {
	var prevColor T
	changes := 0
	for i, currColor := range colorSequence {
		if i > 0 && prevColor != currColor {
			changes += 1
		}
		prevColor = currColor
	}
	return changes
}
