package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Value = int

func BooleanDomain() []Value {
	return []Value{1, 0}
}

func MapDomain[T any](items []T) []Value {
	return fn.NumRange(0, len(items))
}

func IndexDomain(numItems int) []Value {
	return fn.NumRange(0, numItems)
}

func RangeDomain(first, last int) []Value {
	return fn.NumRange(first, last+1)
}
