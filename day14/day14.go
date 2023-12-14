package day14

const (
	ROUND  int = int('O')
	SQUARE int = int('#')
	EMPTY  int = int('.')
)

func Part1(lines []string) int {
	m := len(lines)
	n := len(lines[0])
	g := NewGrid(m, n)

	for r, line := range lines {
		for c, val := range line {
			g.grid[r][c] = int(val)
		}
	}

	for r := 1; r < m; r++ {
		for c := 0; c < n; c++ {
			rollUp(g, r, c)
		}
	}
	// fmt.Println(g)

	sum := 0
	for r, row := range g.grid {
		for _, val := range row {
			if val == ROUND {
				sum += (m - r)
			}
		}
	}

	return sum
}

func rollUp(g *Grid, r, c int) {
	for i := r; i > 0; i-- {
		if g.grid[i][c] != ROUND {
			return
		}
		if g.grid[i-1][c] == EMPTY {
			g.grid[i-1][c] = ROUND
			g.grid[i][c] = EMPTY
		} else {
			break
		}
	}
}
