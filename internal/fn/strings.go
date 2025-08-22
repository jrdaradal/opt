package fn

import "strings"

type stringable interface {
	String() string
}

func ToString[T stringable](item T) string {
	return item.String()
}

func Wrap(items []string) string {
	return "{ " + strings.Join(items, ", ") + " }"
}
