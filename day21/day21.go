package day21

import (
	"aoc2023/common"
	"fmt"
)

func Part1(lines []string) int {
	g, start := NewGrid(lines)
	m := len(g.grid)
	n := len(g.grid[0])

	// for r, row := range g.grid {
	// 	for c := range row {
	// 		if r%2 == 0 && c%2 == 0 {
	// 			g.grid[r][c] = 0
	// 		}
	// 		if r%2 == 1 && c%2 == 1 {
	// 			g.grid[r][c] = 0
	// 		}
	// 	}
	// }

	// T := 0
	// A := 0
	// for r, row := range g.grid {
	// 	for c, val := range row {
	// 		if val == 1 {
	// 			T++
	// 		}
	// 		if r+c <= 64 && val == 1 {
	// 			A++
	// 		}
	// 	}
	// }
	// fmt.Println("T", T)
	// fmt.Println("A", A)

	nrSteps := 65 + 131*2
	// nrSteps := 5 + 11*2
	visitNext := map[Coord]bool{start: true}
	for s := 1; s <= nrSteps; s++ {
		toVisit := visitNext
		visitNext = map[Coord]bool{}
		for pos := range toVisit {
			neighbors := []Coord{NewCoord(pos.r, pos.c-1), NewCoord(pos.r, pos.c+1), NewCoord(pos.r-1, pos.c), NewCoord(pos.r+1, pos.c)}
			for _, nei := range neighbors {
				if g.IsInsideAndEmptyCoord(modCoord(nei, m, n)) {
					visitNext[nei] = true
				}
			}
		}
	}
	printWithOsAfterNSteps(g, &visitNext, nrSteps)

	return len(visitNext)
}

func modCoord(coord Coord, m, n int) Coord {
	return NewCoord(((coord.r+m)%m+m)%m, ((coord.c+n)%n+n)%n)
}

func printWithOs(g *Grid, visit *map[Coord]bool) {
	for r := range g.grid {
		for c, val := range g.grid[r] {
			if (*visit)[NewCoord(r, c)] {
				fmt.Print("O")
			} else if val == EMPTY {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func printWithOsAfterNSteps(g *Grid, visit *map[Coord]bool, n int) {
	R := len(g.grid)
	// r := (R - 1) / 2

	// max := (n - r) / R
	// min := -max
	min, max := 0, 0
	s := 0
	for r := min * R; r < (1+max)*R; r++ {
		// for r := 0; r < R; r++ {
		for c := min * R; c < (1+max)*R; c++ {
			// for c := R; c < 2*R; c++ {
			pos := NewCoord(r, c)
			modPos := modCoord(pos, R, R)
			if (*visit)[pos] {
				fmt.Print("O")
				dx := common.IntAbs(modPos.c - 65)
				dy := common.IntAbs(modPos.r - 65)
				if dx+dy <= 65 {
					s++
				}
			} else if g.grid[modPos.r][modPos.c] == EMPTY {
				fmt.Print(".")
			} else {
				if (r+c)%2 == 0 {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("s", s)
}
