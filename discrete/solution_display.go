package discrete

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/jrdaradal/opt/internal/fn"
)

type SolutionDisplay func(*Solution) string

func DisplayValues[V any](p *Problem, valueMap []V) SolutionDisplay {
	return func(solution *Solution) string {
		output := fn.Map(p.Variables, func(x Variable) string {
			value := solution.Map[x]
			if valueMap == nil {
				return fmt.Sprintf("%d", value)
			} else {
				return fmt.Sprintf("%v", valueMap[value])
			}
		})
		return strings.Join(output, " ")
	}
}

func DisplaySubset[T cmp.Ordered](variableMap []T) SolutionDisplay {
	return func(solution *Solution) string {
		subset := fn.Map(solution.AsSubset(), func(x Variable) T {
			return variableMap[x]
		})
		slices.SortFunc(subset, cmp.Compare)
		output := fn.Map(subset, func(item T) string {
			return fmt.Sprintf("%v", item)
		})
		return fn.Wrap(output)
	}
}

func DisplayPartitions[T any](domain []Value, variableMap []T) SolutionDisplay {
	return func(solution *Solution) string {
		groups := partitionStrings(solution, domain, variableMap)
		outputs := fn.Map(groups, func(group []string) string {
			sort.Strings(group)
			return fn.Wrap(group)
		})
		return strings.Join(outputs, " ")
	}
}

func DisplayMap[T any, V any](p *Problem, variableMap []T, valueMap []V) SolutionDisplay {
	return func(solution *Solution) string {
		output := fn.Map(p.Variables, func(x Variable) string {
			value := solution.Map[x]
			var text1, text2 string
			if variableMap == nil {
				text1 = fmt.Sprintf("%d", x)
			} else {
				text1 = fmt.Sprintf("%v", variableMap[x])
			}
			if valueMap == nil {
				text2 = fmt.Sprintf("%d", value)
			} else {
				text2 = fmt.Sprintf("%v", valueMap[value])
			}
			return text1 + " = " + text2
		})
		return fn.Wrap(output)
	}
}
