package day07

import (
	"aoc2023/common"
	"fmt"
	"sort"
	"strings"
)

func Part1(lines []string) int {
	N := len(lines)

	input := make([][]int, N)
	for r := range input {
		input[r] = make([]int, 7)
	}

	for r, line := range lines {
		tokens := strings.Split(line, " ")
		// fmt.Println(tokens)
		input[r][5] = common.StringToInt(tokens[1])
		for i, run := range tokens[0] {
			input[r][i] = CardToInt(run)
		}
		input[r][6] = CardsToType(input[r][0:5])
	}
	// fmt.Println(input)

	sort.Slice(input, func(i, j int) bool {
		// typi := CardsToType(input[i][0:5])
		// typj := CardsToType(input[j][0:5])
		if input[i][6] < input[j][6] {
			return true
		}
		if input[i][6] > input[j][6] {
			return false
		}
		for ind := 0; ind < 5; ind++ {
			if input[i][ind] < input[j][ind] {
				return true
			}
			if input[i][ind] > input[j][ind] {
				return false
			}
		}
		return false // equal!
	})
	// fmt.Println(input)

	sum := 0
	for i := range input {
		fmt.Printf("%3d %v\n", i, input[i][0:7])
		sum += (i + 1) * input[i][5]
	}

	return sum
}
