package ds

import "strings"

type Subsets struct {
	Universal []string
	Names     []string
	Subsets   [][]string
}

func NewSubsets(universal string, subsetLines []string) *Subsets {
	numSubsets := len(subsetLines)
	names := make([]string, numSubsets)
	subsets := make([][]string, numSubsets)
	for i, line := range subsetLines {
		parts := strings.Split(line, ":")
		names[i] = strings.TrimSpace(parts[0])
		subsets[i] = strings.Fields(strings.TrimSpace(parts[1]))
	}
	return &Subsets{
		Universal: strings.Fields(universal),
		Names:     names,
		Subsets:   subsets,
	}
}
