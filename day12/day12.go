package day12

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

const (
	OPERATIONAL int = int('.')
	DAMAGED     int = int('#')
	UNKNOWN     int = int('?')
)

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		count := countArrangements(line, false)
		// fmt.Println(count, line)
		sum += count
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		count := countArrangements(line, true)
		// fmt.Println(count, line)
		sum += count
	}
	return sum
}

func countArrangements(lineStr string, unfold bool) int {
	tokens := strings.Split(lineStr, " ")
	if len(tokens) != 2 {
		panic("wrong split")
	}
	line := tokens[0]

	tokens = strings.Split(tokens[1], ",")
	rangeLengths := make([]int, len(tokens))
	for i, t := range tokens {
		rangeLengths[i] = common.StringToInt(t)
	}
	// fmt.Println(line, rangeLengths)
	// fmt.Println(len(line), compressedLength(rangeLengths))

	if unfold {
		line = fmt.Sprintf("%s?%s?%s?%s?%s", line, line, line, line, line)
		newRangeLengths := make([]int, len(rangeLengths)*5)
		for i := range newRangeLengths {
			newRangeLengths[i] = rangeLengths[i%len(rangeLengths)]
		}
		rangeLengths = newRangeLengths
	}

	cache := NewCache()
	return rec(line, rangeLengths, cache)
}

func rec(line string, rangeLengths []int, cache *Cache) int {
	cacheValue, exists := cache.Get(line, rangeLengths)
	if exists {
		return cacheValue
	}

	if len(rangeLengths) == 0 {
		panic("no ranges")
	}
	this := rangeLengths[0]
	rest := rangeLengths[1:]
	count := 0

	// Push all remaining ranges (rest) to the right, and compute the space
	// remaining for the current range to move from 0 (start of current line)
	// until 'lastGoodIndex'.
	compressedRest := compressedLength(rest)
	if compressedRest > 0 {
		compressedRest += 1 // leave a separator
	}
	lastGoodIndex := len(line) - compressedRest - this
	// For each possible start position of the current range ...
	for i := 0; i <= lastGoodIndex; i++ {
		// ... check if the range would fit while not contradicting any previous #
		if oneRangeFits(line, i, this) && allNonDamaged(line, 0, i-1) {
			if len(rest) > 0 {
				// This range has a valid position, count the amount of arrangements for
				// the remaining line after removing the current range, at the current position.
				count += rec(line[i+this+1:], rest, cache)
			} else {
				// count the solution, iff all the remaining positions can be operational (aka non-damaged)
				if allNonDamaged(line, i+this, len(line)-1) {
					count++
				}
			}
		}
	}
	cache.Put(line, rangeLengths, count)
	return count
}

func allNonDamaged(line string, from, to int) bool {
	if from > to {
		return true
	}
	for i := from; i <= to; i++ {
		if line[i] == byte(DAMAGED) {
			return false
		}
	}
	return true
}

func compressedLength(rangeLengths []int) int {
	if len(rangeLengths) == 0 {
		return 0
	}
	sum := 0
	for _, r := range rangeLengths {
		sum += r
	}
	sum += (len(rangeLengths) - 1)
	return sum
}

func oneRangeFits(line string, start, size int) bool {
	if !NewRange(start, size).fits(line) {
		return false
	}
	if start > 0 && line[start-1] == byte(DAMAGED) {
		return false
	}
	if start+size <= len(line)-1 && line[start+size] == byte(DAMAGED) {
		return false
	}
	return true
}
