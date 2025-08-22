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

func DisplaySubset[T cmp.Ordered](variableMap []T) SolutionDisplay {
	return func(solution *Solution) string {
		subset := fn.Map(solution.AsSubset(), func(x int) T {
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
