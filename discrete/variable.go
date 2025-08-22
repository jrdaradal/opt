package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Variable = int

func Variables[T any](items []T) []Variable {
	return fn.NumRange(0, len(items))
}

func RangeVariables(first, last int) []Variable {
	return fn.NumRange(first, last+1)
}
