package discrete

import (
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
