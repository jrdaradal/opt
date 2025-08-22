package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Solution struct {
	Map map[Variable]Value
	Score
}

func (s Solution) Values() []Value {
	return fn.MapValues(s.Map)
}

// List of values ordered by variable order
func (s Solution) Tuple(p *Problem) []Value {
	tuple := make([]Value, len(p.Variables))
	for i, variable := range p.Variables {
		tuple[i] = s.Map[variable]
	}
	return tuple
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

func (s Solution) AsPartitions(domain []Value) [][]Variable {
	groups := make(map[Value][]Variable)
	for _, value := range domain {
		groups[value] = make([]Variable, 0)
	}
	for variable, value := range s.Map {
		groups[value] = append(groups[value], variable)
	}
	partitions := make([][]Variable, len(domain))
	for i, value := range domain {
		partitions[i] = groups[value]
	}
	return partitions
}

func (s Solution) PartitionSums(domain []Value, valueMap []int) []int {
	return fn.Map(s.AsPartitions(domain), func(partition []int) int {
		return fn.SumValues(partition, valueMap)
	})
}

func partitionStrings[T any](solution *Solution, domain []Value, variableMap []T) [][]string {
	return fn.Map(solution.AsPartitions(domain), func(partition []int) []string {
		return fn.TranslateString(partition, variableMap)
	})
}
