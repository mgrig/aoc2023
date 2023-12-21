package day21

import "fmt"

func Part1(lines []string) int {
	g, start := NewGrid(lines)

	nrSteps := 64
	visitNext := map[Coord]bool{start: true}
	for s := 1; s <= nrSteps; s++ {
		toVisit := visitNext
		visitNext = map[Coord]bool{}
		for pos := range toVisit {
			neighbors := []Coord{NewCoord(pos.r, pos.c-1), NewCoord(pos.r, pos.c+1), NewCoord(pos.r-1, pos.c), NewCoord(pos.r+1, pos.c)}
			for _, n := range neighbors {
				if g.IsInsideAndEmptyCoord(n) {
					visitNext[n] = true
				}
			}
		}
		// fmt.Println(len(visitNext), visitNext)
		// fmt.Println(s)
		// printWithOs(g, &visitNext)
	}

	return len(visitNext)
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
