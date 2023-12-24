package day24

type Line struct {
	x, y   int
	dx, dy int
	x2, y2 int
}

func NewLine(x, y, dx, dy int) Line {
	return Line{
		x: x, y: y,
		dx: dx, dy: dy,
		x2: x + dx, y2: y + dy,
	}
}
