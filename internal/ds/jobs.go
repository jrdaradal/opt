package ds

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/jrdaradal/opt/internal/fn"
)

const taskGlue string = ":"

type TimeRange [2]int

type Job struct {
	ID    int
	Tasks []*Task
}

type Task struct {
	ID       string
	JobID    int
	Machine  string
	Duration int
}

func NewJob(line string, jobID int) *Job {
	job := &Job{
		ID:    jobID,
		Tasks: make([]*Task, 0),
	}
	for taskID, text := range strings.Fields(line) {
		task := NewTask(text, jobID, taskID)
		job.Tasks = append(job.Tasks, task)
	}
	return job
}

func NewTask(text string, jobID int, taskID int) *Task {
	parts := fn.CleanSplit(text, taskGlue)
	if len(parts) != 2 {
		return nil
	}
	return &Task{
		ID:       fmt.Sprintf("J%d_T%d", jobID, taskID),
		JobID:    jobID,
		Machine:  parts[0],
		Duration: fn.ParseInt(parts[1]),
	}
}

func (j Job) TotalDuration() int {
	return fn.Sum(fn.Map(j.Tasks, taskDuration))
}

func (j Job) TaskMargins(taskID int) (int, int) {
	before := fn.Sum(fn.Map(j.Tasks[:taskID], taskDuration))
	after := fn.Sum(fn.Map(j.Tasks[taskID+1:], taskDuration))
	return before, after
}

func taskDuration(task *Task) int {
	return task.Duration
}

func SortByTimeRange(a, b TimeRange) int {
	return cmp.Compare(a[0], b[0])
}

func (t TimeRange) Tuple() (int, int) {
	return t[0], t[1]
}
