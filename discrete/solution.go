package discrete

import "github.com/jrdaradal/opt/internal/fn"

type Solution struct {
	Map map[Variable]Value
	Score
}

func (s Solution) Values() []Value {
	return fn.MapValues(s.Map)
}
