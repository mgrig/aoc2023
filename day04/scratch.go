package day04

import (
	"aoc2023/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Part1(lines []string) int {
	re := regexp.MustCompile(`Card +\d+: (.*) \| (.*)`)

	sum := 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)

		winning := stringToIntSet(parts[1])
		// fmt.Println(winning)

		have := stringToIntArray(parts[2])
		// fmt.Println(have)

		count := 0
		for _, h := range have {
			if contains(&winning, h) {
				count++
			}
		}
		// fmt.Println(count, int(math.Pow(2.0, float64(count-1)))
		points := int(math.Pow(2.0, float64(count-1)))

		sum += points
	}

	return sum
}

func Part2(lines []string) int {
	re := regexp.MustCompile(`Card +(\d+): (.*) \| (.*)`)

	nrOwned := make(map[int]int, len(lines))
	for i := 1; i <= len(lines); i++ {
		nrOwned[i] = 1
	}

	sum := 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		cardNr := common.StringToInt(parts[1])
		countWins := countWins(parts[2], parts[3])

		for i := 1; i <= countWins; i++ {
			nrOwned[cardNr+i] += nrOwned[cardNr]
		}
		sum += nrOwned[cardNr]
	}
	fmt.Println(nrOwned)

	return sum
}

func countWins(winningStr, haveStr string) int {
	winning := stringToIntSet(winningStr)
	have := stringToIntArray(haveStr)

	count := 0
	for _, h := range have {
		if contains(&winning, h) {
			count++
		}
	}

	return count
}

func contains(set *map[int]bool, value int) bool {
	_, exists := (*set)[value]
	return exists
}

func stringToIntSet(str string) map[int]bool {
	tokens := strings.Split(strings.Trim(str, " "), " ")
	ret := make(map[int]bool, 0)
	for _, token := range tokens {
		if token == "" {
			continue
		}
		ret[common.StringToInt(token)] = true
	}
	return ret
}

func stringToIntArray(str string) []int {
	tokens := strings.Split(strings.Trim(str, " "), " ")
	ret := make([]int, 0)
	for _, token := range tokens {
		if token == "" {
			continue
		}
		ret = append(ret, common.StringToInt(token))
	}
	return ret
}
