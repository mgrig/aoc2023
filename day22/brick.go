package day22

import (
	"aoc2023/common"
	"fmt"
)

type Brick struct {
	from, to Coord
}

func NewBrick(from, to Coord) Brick {
	return Brick{
		from: from,
		to:   to,
	}
}

func (b Brick) MinZ() int {
	return common.IntMin(b.from.z, b.to.z)
}

func (b Brick) MaxZ() int {
	return common.IntMax(b.from.z, b.to.z)
}

func (b Brick) OuterBox() (minx, maxx, miny, maxy, minz, maxz int) {
	return common.IntMin(b.from.x, b.to.x), common.IntMax(b.from.x, b.to.x),
		common.IntMin(b.from.y, b.to.y), common.IntMax(b.from.y, b.to.y),
		common.IntMin(b.from.z, b.to.z), common.IntMax(b.from.z, b.to.z)
}

func (b Brick) String() string {
	return fmt.Sprintf("%s - %s", b.from, b.to)
}

// returns with z = 0
func (b Brick) HorizProjection() []Coord {
	ret := []Coord{}

	minx, maxx, miny, maxy, _, _ := b.OuterBox()
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			ret = append(ret, NewCoord(x, y, 0))
		}
	}

	return ret
}

func (b *Brick) Lower(n int) {
	b.from.z -= n
	b.to.z -= n
}

func (b Brick) AllCoords() []Coord {
	ret := []Coord{}

	minx, maxx, miny, maxy, minz, maxz := b.OuterBox()
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			for z := minz; z <= maxz; z++ {
				ret = append(ret, NewCoord(x, y, z))
			}
		}
	}

	return ret
}
