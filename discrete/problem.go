package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Problem struct {
	Name        string
	Variables   []Variable
	Domain      map[Variable][]Value
	Constraints []Constraint
	Goal
	ObjectiveFunc
	SolutionCore
	SolutionDisplay
}

func NewProblem(name string) *Problem {
	return &Problem{
		Name:         name,
		Variables:    make([]Variable, 0),
		Domain:       make(map[Variable][]Value),
		Constraints:  make([]Constraint, 0),
		SolutionCore: nil,
	}
}

func (p *Problem) AddConstraint(c Constraint) {
	p.Constraints = append(p.Constraints, c)
}

func (p *Problem) AddGlobalConstraint(test ConstraintFunc) {
	p.AddConstraint(NewGlobalConstraint(p, test))
}

func (p Problem) SolutionSpace() int {
	size := 1
	for _, domain := range p.Domain {
		size *= len(domain)
	}
	return size
}

func (p Problem) IsSatisfied(solution *Solution) bool {
	return fn.All(p.Constraints, func(c Constraint) bool {
		return c.IsSatisfied(solution)
	})
}

func (p Problem) IsSatisfactionProblem() bool {
	return p.Goal == SATISFY
}

func (p Problem) IsOptimizationProblem() bool {
	return p.Goal == MINIMIZE || p.Goal == MAXIMIZE
}
