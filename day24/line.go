package day24

type Line struct {
	start Coord
	dir   Coord
}

func NewLine(x, y, z, dx, dy, dz int) Line {
	return Line{
		start: NewCoord(x, y, z),
		dir:   NewCoord(dx, dy, dz),
	}
}

func (l Line) PositionAt(time int) Coord {
	return NewCoord(l.start.x+time*l.dir.x, l.start.y+time*l.dir.y, l.start.z+time*l.dir.z)
}
