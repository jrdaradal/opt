package fn

import (
	"strconv"
	"strings"
)

func NumRange(start, end int) []int {
	numbers := make([]int, end-start)
	for n := start; n < end; n++ {
		numbers[n-start] = n
	}
	return numbers
}

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func Abs(x int) int {
	return Ternary(x < 0, -x, x)
}

func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	return Ternary(err == nil, number, 0)
}

func ParseInt(value string) int {
	number, err := strconv.Atoi(value)
	return Ternary(err == nil, number, 0)
}
