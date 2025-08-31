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
	ID    string
	Tasks []*Task
}

type Task struct {
	ID       string
	JobID    string
	Machine  string
	Duration int
}

type SlotSched struct {
	Start int
	End   int
	Name  string
}

func NewJob(line string, jobID string) *Job {
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

func NewTask(text string, jobID string, taskID int) *Task {
	parts := fn.CleanSplit(text, taskGlue)
	if len(parts) != 2 {
		return nil
	}
	return &Task{
		ID:       fmt.Sprintf("J%s_T%d", jobID, taskID),
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

func SortByStartTime(a, b TimeRange) int {
	return cmp.Compare(a[0], b[0])
}

func SortBySchedStart(a, b SlotSched) int {
	return cmp.Compare(a.Start, b.Start)
}

func (t TimeRange) Tuple() (int, int) {
	return t[0], t[1]
}

func TaskString(machine string, duration string) string {
	return fmt.Sprintf("%s%s%s", machine, taskGlue, duration)
}

func (j Job) String() string {
	return j.ID
}
