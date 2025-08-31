package ds

type StrInt struct {
	Str string
	Int int
}

func (s StrInt) Tuple() (string, int) {
	return s.Str, s.Int
}
