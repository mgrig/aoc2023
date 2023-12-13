package day13

import (
	"aoc2023/common"
)

func Part1(lines []string) int {
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
		h := hist(g.grid)
		splitIndex, found := searchReflectionLine(h, len(g.grid))
		if found {
			sum += 100 * splitIndex
		} else {
			ht := hist(g.Transpose().grid)
			splitIndex, found = searchReflectionLine(ht, len(g.grid[0]))
			if !found {
				panic("nothing found")
			}
			sum += splitIndex
		}
	}

	return sum
}

func histToSparse(h map[string][]int, nrIndexes int) *Sparse {
	ret := NewSparse(nrIndexes)
	for _, hval := range h {
		for _, iVal := range hval {
			ret.Set(iVal, hval)
		}
	}
	return ret
}

func hist(lines []string) map[string][]int {
	ret := make(map[string][]int, 0)
	for i, line := range lines {
		_, exists := ret[line]
		if !exists {
			ret[line] = make([]int, 0)
		}
		ret[line] = append(ret[line], i)
	}
	return ret
}

func searchReflectionLine(h map[string][]int, nrIndexes int) (index int, found bool) {
	// one of the 2 ends must belong to the reflected surface. we try both.
	s := (*histToSparse(h, nrIndexes)).mat
	canStartLeft := !(len(s[0]) == 1 || len(s[1]) == 1)
	canStartRight := !(len(s[len(s)-1]) == 1 || len(s[len(s)-2]) == 1)
	if !canStartLeft && !canStartRight {
		return -1, false
	}

	if canStartLeft {
		splits := possibleSplitIndexes(s[0])
		for _, split := range splits {
			to := 2*split - 1
			if to > len(s)-1 {
				continue
			}
			if checkSplit(split, 0, to, &s) { //TODO could there be more solutions?!
				return split, true
			}
		}
	}

	// can start right
	splits := possibleSplitIndexes(s[len(s)-1])
	for _, split := range splits {
		lenSplit := len(s) - split
		from := len(s) - 2*lenSplit
		if from < 0 {
			continue
		}
		if checkSplit(split, from, len(s)-1, &s) { //TODO could there be more solutions?!
			return split, true
		}
	}

	return -1, false
}

func checkSplit(split int, from, to int, sparse *[][]int) bool {
	for i := from; i <= to; i++ {
		if !sliceContains(possibleSplitIndexes((*sparse)[i]), split) {
			return false
		}
	}
	return true
}

func sliceContains(slice []int, value int) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func possibleSplitIndexes(histIndexes []int) []int {
	ret := make([]int, 0)
	if len(histIndexes) < 2 {
		return ret
	}

	for i := 0; i < len(histIndexes)-1; i++ {
		iVal := histIndexes[i]
		for j := i + 1; j < len(histIndexes); j++ {
			jVal := histIndexes[j]
			if common.IntAbs(iVal-jVal)%2 == 1 {
				ret = append(ret, (iVal+jVal+1)/2)
			}
		}
	}

	return ret
}
