package ds

type Coords [2]int

func (c Coords) Tuple() (int, int) {
	return c[0], c[1]
}
