package day23

type Coord struct {
	r, c int
}

func NewCoord(r, c int) Coord {
	return Coord{r: r, c: c}
}
