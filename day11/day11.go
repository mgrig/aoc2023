package day11

import (
	"aoc2023/common"
)

func Part1(lines []string, expansion int) int {
	m := len(lines)
	n := len(lines[0])

	emptyRows := make(map[int]bool, m)
	for r := 0; r < m; r++ {
		emptyRows[r] = true
	}

	emptyCols := make(map[int]bool, n)
	for c := 0; c < m; c++ {
		emptyCols[c] = true
	}

	galaxies := make([]Coord, 0)
	for r := range lines {
		for c, val := range lines[r] {
			if val == '#' {
				delete(emptyRows, r)
				delete(emptyCols, c)
				galaxies = append(galaxies, NewCoord(r, c))
			}
		}
	}
	// fmt.Println(len(galaxies), emptyRows, emptyCols)

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			from := galaxies[i]
			to := galaxies[j]
			d := dist(from, to, &emptyRows, &emptyCols, expansion)
			// fmt.Println(i, j, "d=", d)
			sum += d
		}
	}

	return sum
}

func dist(from, to Coord, emptyRows, emptyCols *map[int]bool, expansion int) int {
	dx := common.IntAbs(to.c - from.c)
	if dx > 1 {
		min := common.IntMin(from.c, to.c)
		max := common.IntMax(from.c, to.c)
		dx += (expansion - 1) * countEmptyBetween(min+1, max-1, emptyCols)
	}

	dy := common.IntAbs(to.r - from.r)
	if dy > 1 {
		min := common.IntMin(from.r, to.r)
		max := common.IntMax(from.r, to.r)
		dy += (expansion - 1) * countEmptyBetween(min+1, max-1, emptyRows)
	}

	return dx + dy
}

func countEmptyBetween(from, to int, empty *map[int]bool) int {
	count := 0
	for k := range *empty {
		if from <= k && k <= to {
			count++
		}
	}
	return count
}
