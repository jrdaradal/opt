package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Solution struct {
	Map map[Variable]Value
	Score
}

func (s Solution) Values() []Value {
	return fn.MapValues(s.Map)
}

// Assumes BooleanDomain {0,1}
func (s Solution) AsSubset() []Variable {
	subset := make([]Variable, 0)
	for variable, value := range s.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// Assumes BooleanDomain {0,1}
func (s Solution) SubsetCount() int {
	return fn.Sum(fn.MapValues(s.Map))
}
