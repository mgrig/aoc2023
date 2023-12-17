package day17

import "fmt"

type Grid struct {
	grid [][]int
}

func NewGrid(n int) *Grid {
	grid := make([][]int, n)
	for r := range grid {
		grid[r] = make([]int, n)
	}
	return &Grid{
		grid: grid,
	}
}

func (g *Grid) String() string {
	ret := ""
	for r := range g.grid {
		for _, val := range g.grid[r] {
			ret += fmt.Sprintf("%d", val)
		}
		ret += fmt.Sprintln()
	}
	return ret
}

func (g *Grid) Inside(r, c int) bool {
	n := len(g.grid)
	return r >= 0 && c >= 0 && r < n && c < n
}
