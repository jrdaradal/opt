package discrete

import "github.com/jrdaradal/opt/internal/fn"

type GlobalConstraint struct {
	BaseConstraint
}

func NewGlobalConstraint(p *Problem, test ConstraintFunc) GlobalConstraint {
	c := GlobalConstraint{}
	c.Variables = p.Variables
	c.Test = test
	c.Penalty = fn.Ternary(p.Goal == MAXIMIZE, -HARD_PENALTY, HARD_PENALTY)
	return c
}

func (c GlobalConstraint) ComputePenalty(solution *Solution) Penalty {
	if !c.IsSatisfied(solution) {
		return c.Penalty
	}
	return 0
}
