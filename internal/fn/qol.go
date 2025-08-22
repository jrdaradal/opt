package fn

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
