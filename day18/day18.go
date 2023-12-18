package day18

import (
	"aoc2023/common"
	"fmt"
	"regexp"
)

const (
	U int = 0
	D int = 1
	L int = 2
	R int = 3

	EMPTY int = 0
	DIG   int = 1
)

func Part1(lines []string) int {
	reDig := regexp.MustCompile(`(.+?) (\d+?) \(.*\)`)

	r, c := 0, 0
	minr, minc, maxr, maxc := 0, 0, 0, 0
	coords := make([]Coord, 0)
	for _, line := range lines {
		tokens := reDig.FindStringSubmatch(line)
		if len(tokens) != 3 {
			panic("wrong regex match")
		}

		dir := parseUDLR(tokens[1])
		count := common.StringToInt(tokens[2])

		switch dir {
		case U:
			newr := r - count
			for ir := r; ir >= newr; ir-- {
				coords = append(coords, NewCoord(ir, c))
			}
			r = newr
		case D:
			// r += count
			newr := r + count
			for ir := r; ir <= newr; ir++ {
				coords = append(coords, NewCoord(ir, c))
			}
			r = newr
		case L:
			// c -= count
			newc := c - count
			for ic := c; ic >= newc; ic-- {
				coords = append(coords, NewCoord(r, ic))
			}
			c = newc
		case R:
			// c += count
			newc := c + count
			for ic := c; ic <= newc; ic++ {
				coords = append(coords, NewCoord(r, ic))
			}
			c = newc
		}
		minr = common.IntMin(minr, r)
		maxr = common.IntMax(maxr, r)
		minc = common.IntMin(minc, c)
		maxc = common.IntMax(maxc, c)
	}
	fmt.Println(minr, minc, " - ", maxr, maxc)

	m := maxr - minr + 1
	n := maxc - minc + 1
	g := NewGrid(m, n)
	for _, coord := range coords {
		g.grid[coord.r-minr][coord.c-minc] = DIG
	}
	// fmt.Println(g)

	// sum := 0
	// for _, row := range g.grid {
	// 	for _, val := range row {
	// 		if val == DIG {
	// 			sum++
	// 		}
	// 	}
	// }

	corners := NewBoolGrid(m+1, n+1)
	toPropagate := []Coord{NewCoord(0, 0)}
	for len(toPropagate) > 0 {
		propagateOuter(corners, g, &toPropagate)
	}
	// fmt.Println(corners)

	// count inner cells
	sum := 0
	for r := range g.grid {
		for c := range g.grid[r] {
			if !corners.GetAt(r, c) || !corners.GetAt(r, c+1) || !corners.GetAt(r+1, c+1) || !corners.GetAt(r+1, c) {
				sum++
			}
		}
	}

	return sum
}

func propagateOuter(corners *BoolGrid, g *Grid, toPropagate *[]Coord) {
	pos := (*toPropagate)[0]
	*toPropagate = (*toPropagate)[1:]

	if !corners.InsideCoord(pos) || corners.GetAt(pos.r, pos.c) {
		return
	}
	corners.grid[pos.r][pos.c] = true

	right := NewCoord(pos.r, pos.c+1)
	if corners.InsideCoord(right) && !( /*blocked*/ g.GetAt(pos.r, pos.c) == DIG && g.GetAt(pos.r-1, pos.c) == DIG) {
		*toPropagate = append(*toPropagate, right)
	}

	left := NewCoord(pos.r, pos.c-1)
	if corners.InsideCoord(left) && !( /*blocked*/ g.GetAt(pos.r, pos.c-1) == DIG && g.GetAt(pos.r-1, pos.c-1) == DIG) {
		*toPropagate = append(*toPropagate, left)
	}

	up := NewCoord(pos.r-1, pos.c)
	if corners.InsideCoord(up) && !( /*blocked*/ g.GetAt(pos.r-1, pos.c-1) == DIG && g.GetAt(pos.r-1, pos.c) == DIG) {
		*toPropagate = append(*toPropagate, up)
	}

	down := NewCoord(pos.r+1, pos.c)
	if corners.InsideCoord(down) && !( /*blocked*/ g.GetAt(pos.r, pos.c-1) == DIG && g.GetAt(pos.r, pos.c) == DIG) {
		*toPropagate = append(*toPropagate, down)
	}
}

func parseUDLR(s string) int {
	switch s {
	case "U":
		return U
	case "D":
		return D
	case "L":
		return L
	case "R":
		return R
	default:
		panic("wrong UDLR")
	}
}
