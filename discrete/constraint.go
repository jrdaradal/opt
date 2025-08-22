package discrete

import "math"

type Penalty = float64

var HARD_PENALTY Penalty = math.Inf(1)

type Constraint interface {
	IsSatisfied(*Solution) bool
	ComputePenalty(*Solution) Penalty
}

type ConstraintFunc func(*Solution) bool

type BaseConstraint struct {
	Penalty
	Variables []Variable
	Test      ConstraintFunc
	// TODO: Add PartialTest for Solvers with PartialSolution
}

func (c BaseConstraint) IsSatisfied(solution *Solution) bool {
	return c.Test(solution)
}
