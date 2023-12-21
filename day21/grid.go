package day21

const (
	EMPTY int = 0
	ROCK  int = 1
)

type Grid struct {
	grid [][]int
}

func NewGrid(lines []string) (grid *Grid, start Coord) {
	g := make([][]int, len(lines))

	var startr, startc int
	for r, row := range lines {
		g[r] = make([]int, len(lines[0]))
		for c, val := range row {
			if int(val) == int('S') {
				g[r][c] = EMPTY
				startr = r
				startc = c
			} else if int(val) == int('.') {
				g[r][c] = EMPTY
			} else {
				g[r][c] = ROCK
			}
		}
	}

	return &Grid{
		grid: g,
	}, NewCoord(startr, startc)
}

func (g *Grid) Inside(r, c int) bool {
	m := len(g.grid)
	n := len(g.grid[0])
	return r >= 0 && r < m && c >= 0 && c < n
}

func (g *Grid) IsInsideAndEmpty(r, c int) bool {
	if !g.Inside(r, c) {
		return false
	}

	return g.grid[r][c] == EMPTY
}

func (g *Grid) IsInsideAndEmptyCoord(coord Coord) bool {
	return g.IsInsideAndEmpty(coord.r, coord.c)
}
