package fn

import (
	"strings"
)

type stringable interface {
	String() string
}

func ToString[T stringable](item T) string {
	return item.String()
}

func Wrap(items []string) string {
	return "{ " + strings.Join(items, ", ") + " }"
}

func CleanSplit(text string, sep string) []string {
	parts := strings.Split(text, sep)
	return Map(parts, strings.TrimSpace)
}
