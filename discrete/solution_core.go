package discrete

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/jrdaradal/opt/internal/fn"
)

type SolutionCore func(*Solution) string

func SortedPartition[T any](domain []Value, variableMap []T) SolutionCore {
	return func(solution *Solution) string {
		groups := partitionStrings(solution, domain, variableMap)
		groups = fn.Filter(groups, func(group []string) bool {
			return len(group) > 0
		})
		partitions := fn.Map(groups, func(group []string) string {
			sort.Strings(group)
			return fn.Wrap(group)
		})
		sort.Strings(partitions)
		return strings.Join(partitions, "/")
	}
}

func LookupValueOrder(p *Problem) SolutionCore {
	return func(solution *Solution) string {
		values := solution.Tuple(p)
		core := make([]string, len(values))
		lookup := make(map[Value]string)
		order := 0
		for i, value := range values {
			if _, ok := lookup[value]; !ok {
				lookup[value] = fmt.Sprintf("%d", order)
				order++
			}
			core[i] = lookup[value]
		}
		return strings.Join(core, "")
	}
}

func MirroredSequence[T any](variableMap []T) SolutionCore {
	return func(solution *Solution) string {
		numItems := len(solution.Map)
		sequence := make([]string, numItems)
		for x, idx := range solution.Map {
			sequence[idx] = fmt.Sprintf("%v", variableMap[x])
		}
		first, last := sequence[0], sequence[numItems-1]
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(sequence)
		}
		return strings.Join(sequence, " ")
	}
}
