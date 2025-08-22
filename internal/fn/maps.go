package fn

import (
	"maps"
	"slices"
)

func MapKeys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

func MapValues[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, ok := items[key]
	return ok
}

func Translate[K comparable, V any](items []K, lookup map[K]V) []V {
	items2 := make([]V, len(items))
	for i, item := range items {
		items2[i] = lookup[item]
	}
	return items2
}

func NewCounter[T comparable](items []T) map[T]int {
	count := make(map[T]int)
	for _, item := range items {
		count[item] = 0
	}
	return count
}

func TallyValues[K, V comparable](items map[K]V, values []V) map[V]int {
	tally := make(map[V]int, len(values))
	for _, value := range values {
		tally[value] = 0
	}
	for _, value := range items {
		tally[value] += 1
	}
	return tally
}

func BooleanMap[K comparable](items []K, flag bool) map[K]bool {
	items2 := make(map[K]bool)
	for _, item := range items {
		items2[item] = flag
	}
	return items2
}
