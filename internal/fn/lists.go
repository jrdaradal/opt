package fn

import "fmt"

func All[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

func AllEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

func AllTrue(items []bool) bool {
	for _, ok := range items {
		if !ok {
			return false
		}
	}
	return true
}

func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

func Map[T any, V any](items []T, convert func(T) V) []V {
	items2 := make([]V, len(items))
	for i, item := range items {
		items2[i] = convert(item)
	}
	return items2
}

func MapIndex[T any, V any](items []T, convert func(int, T) V) []V {
	items2 := make([]V, len(items))
	for i, item := range items {
		items2[i] = convert(i, item)
	}
	return items2
}

func Filter[T any](items []T, keep func(T) bool) []T {
	items2 := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			items2 = append(items2, item)
		}
	}
	return items2
}

func TranslateString[T any](indexes []int, translate []T) []string {
	items2 := make([]string, len(indexes))
	for i, idx := range indexes {
		items2[i] = fmt.Sprintf("%v", translate[idx])
	}
	return items2
}

func Sum[T ~int | ~float64](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}
	return sum
}

func SumValues[T ~int | ~float64](indexes []int, valueMap []T) T {
	return Sum(Map(indexes, func(x int) T {
		return valueMap[x]
	}))
}

func CountValue[T comparable](items []T, value T) int {
	count := 0
	for _, item := range items {
		if item == value {
			count += 1
		}
	}
	return count
}
