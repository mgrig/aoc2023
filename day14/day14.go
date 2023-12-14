package day14

import (
	"fmt"
	"hash/fnv"
)

const (
	ROUND  int = int('O')
	SQUARE int = int('#')
	EMPTY  int = int('.')

	UP    int = 0
	LEFT  int = 1
	DOWN  int = 2
	RIGHT int = 3
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

	rollUp(g, NewHistory(), 0)

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

func Part2(lines []string) int {
	m := len(lines)
	n := len(lines[0])
	g := NewGrid(m, n)

	for r, line := range lines {
		for c, val := range line {
			g.grid[r][c] = int(val)
		}
	}
	h := NewHistory()

	var cycle, prevCycle, rest int
	var found bool
	N := 1_000_000_000
	for cycle = 0; cycle < N; cycle++ {
		fmt.Println(cycle)

		// fmt.Printf("%2d U %d -> ", cycle, hash(g.String()))
		rollUp(g, h, cycle)
		// fmt.Println(hash(g.String()))

		// fmt.Printf("%2d L %d -> ", cycle, hash(g.String()))
		rollLeft(g, h, cycle)
		// fmt.Println(hash(g.String()))

		// fmt.Printf("%2d D %d -> ", cycle, hash(g.String()))
		rollDown(g, h, cycle)
		// fmt.Println(hash(g.String()))

		// fmt.Printf("%2d R %d -> ", cycle, hash(g.String()))
		prevCycle, found = rollRight(g, h, cycle)
		// fmt.Println(hash(g.String()))

		if found {
			fmt.Println("cg hit", prevCycle, cycle)
			beforeRepeat := prevCycle + 1
			period := cycle - prevCycle
			rest = (N - beforeRepeat) % period
			cycle = N - rest
			// fmt.Println("new cycle:", cycle)
			break
		}
	}

	for ; cycle < N; cycle++ {
		fmt.Println(cycle)

		fmt.Printf("%2d U %d -> ", cycle, hash(g.String()))
		rollUp(g, h, cycle)
		fmt.Println(hash(g.String()))

		fmt.Printf("%2d L %d -> ", cycle, hash(g.String()))
		rollLeft(g, h, cycle)
		fmt.Println(hash(g.String()))

		fmt.Printf("%2d D %d -> ", cycle, hash(g.String()))
		rollDown(g, h, cycle)
		fmt.Println(hash(g.String()))

		fmt.Printf("%2d R %d -> ", cycle, hash(g.String()))
		rollRight(g, h, cycle)
		fmt.Println(hash(g.String()))
	}

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

func rollUp(g *Grid, h *History, cycle int) (prevCycle int, cgHit bool) {
	prevCycle, found := h.Get(g, UP)
	if !found {
		h.Set(g, UP, cycle)
	}

	m := len(g.grid)
	n := len(g.grid[0])

	for c := 0; c < n; c++ {
		vector := make([]int, m)
		for r := 0; r < m; r++ {
			vector[r] = g.grid[r][c]
		}

		vector = rollRowLeft(vector)

		for r := 0; r < m; r++ {
			g.grid[r][c] = vector[r]
		}
	}
	if found {
		return prevCycle, true
	} else {
		return -1, false
	}
}

func rollDown(g *Grid, h *History, cycle int) (prevCycle int, cgHit bool) {
	prevCycle, found := h.Get(g, DOWN)
	if !found {
		h.Set(g, DOWN, cycle)
	}

	m := len(g.grid)
	n := len(g.grid[0])

	for c := 0; c < n; c++ {
		vector := make([]int, m)
		for r := m - 1; r >= 0; r-- {
			vector[m-1-r] = g.grid[r][c]
		}

		vector = rollRowLeft(vector)

		for r := m - 1; r >= 0; r-- {
			g.grid[r][c] = vector[m-1-r]
		}
	}
	if found {
		return prevCycle, true
	} else {
		return -1, false
	}
}

func rollLeft(g *Grid, h *History, cycle int) (prevCycle int, cgHit bool) {
	prevCycle, found := h.Get(g, LEFT)
	if !found {
		h.Set(g, LEFT, cycle)
	}

	m := len(g.grid)
	n := len(g.grid[0])

	for r := 0; r < m; r++ {
		vector := make([]int, n)
		for c := 0; c < n; c++ {
			vector[c] = g.grid[r][c]
		}

		vector = rollRowLeft(vector)

		for c := 0; c < n; c++ {
			g.grid[r][c] = vector[c]
		}
	}
	if found {
		return prevCycle, true
	} else {
		return -1, false
	}
}

func rollRight(g *Grid, h *History, cycle int) (prevCycle int, cgHit bool) {
	prevCycle, found := h.Get(g, RIGHT)
	if !found {
		h.Set(g, RIGHT, cycle)
	}

	m := len(g.grid)
	n := len(g.grid[0])

	for r := 0; r < m; r++ {
		vector := make([]int, n)
		for c := n - 1; c >= 0; c-- {
			vector[n-1-c] = g.grid[r][c]
		}

		vector = rollRowLeft(vector)

		for c := n - 1; c >= 0; c-- {
			g.grid[r][c] = vector[n-1-c]
		}
	}
	if found {
		return prevCycle, true
	} else {
		return -1, false
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func rollRowLeft(row []int) []int {
	ret := make([]int, len(row))
	copy(ret, row)
	for i := 1; i < len(ret); i++ {
		for j := i; j > 0; j-- {
			if ret[j] != ROUND {
				break
			}
			if ret[j-1] == EMPTY {
				ret[j-1] = ROUND
				ret[j] = EMPTY
			} else {
				break
			}
		}
	}
	return ret
}
