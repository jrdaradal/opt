package main

import (
	"fmt"

	"github.com/jrdaradal/opt/discrete"
	"github.com/jrdaradal/opt/problem"
)

func main() {
	p := problem.CarSequencing("carsequence1")
	solution := discrete.EmptySolution()
	for i := range 10 {
		solution.Assign(i, i)
	}
	t2 := p.Constraints[1]

	fmt.Print(t2.IsSatisfied(solution))
}
