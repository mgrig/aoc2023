package day13

func Part(lines []string, maxSmudges int) int {
	grids := make([]*Grid, 0)
	grid := NewGrid()
	grids = append(grids, grid)
	for _, line := range lines {
		if line != "" {
			grid.grid = append(grid.grid, line)
		} else {
			grid = NewGrid()
			grids = append(grids, grid)
		}
	}

	sum := 0
	for _, g := range grids {
		sum += summarize(g, maxSmudges)
	}

	return sum
}

func summarize(g *Grid, maxSmudges int) int {
	for split := 1; split < len(g.grid); split++ {
		if isValidSplit(g.grid, split, maxSmudges) {
			return 100 * split
		}
	}

	g = g.Transpose()
	for split := 1; split < len(g.grid); split++ {
		if isValidSplit(g.grid, split, maxSmudges) {
			return split
		}
	}

	panic("split not found")
}

func isValidSplit(lines []string, splitIndex int, maxSmudges int) bool {
	n := len(lines)
	smudges := 0 // applied smudges
	for left, right := splitIndex-1, splitIndex; left >= 0 && right <= n-1; left, right = left-1, right+1 {
		delta := diff(lines[left], lines[right])
		if smudges+delta <= maxSmudges {
			if delta == 1 {
				smudges = 1
			}
		} else {
			return false // too many diffs
		}
	}
	return smudges == maxSmudges
}

func diff(line1, line2 string) int {
	if len(line1) != len(line2) {
		panic("wrong lines")
	}
	sum := 0
	for i, c1 := range line1 {
		if byte(c1) != line2[i] {
			sum++
		}
	}
	return sum
}
