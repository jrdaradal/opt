package fn

import (
	"fmt"
	"os"
	"strings"
)

func ProblemData(name string) ([]string, error) {
	path := fmt.Sprintf("data/%s.txt", name)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	lines = Map(lines, strings.TrimSpace)
	lines = Filter(lines, func(line string) bool {
		return !strings.HasPrefix(line, "#") && line != ""
	})
	return lines, nil
}
