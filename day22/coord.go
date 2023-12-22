package day22

import "fmt"

type Coord struct {
	x, y, z int
}

func NewCoord(x, y, z int) Coord {
	return Coord{x: x, y: y, z: z}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}
