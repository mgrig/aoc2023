package day09

import (
	"aoc2023/common"
	"strings"
)

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		intLine := stringToIntArray(line)
		sum += computeLine(intLine)
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		intLine := stringToIntArray(line)
		sum += computeLineLeft(intLine)
	}
	return sum
}

func computeLineLeft(intLine []int) int {
	if allZeros(intLine) {
		return 0
	}

	sum := 0
	sign := 1
	allZeros := false
	for !allZeros {
		sum += sign * intLine[0]
		sign *= -1
		intLine, allZeros = diff(intLine)
	}
	return sum
}

func computeLine(intLine []int) int {
	if allZeros(intLine) {
		return 0
	}

	sum := intLine[len(intLine)-1]
	allZeros := false
	for !allZeros {
		intLine, allZeros = diff(intLine)
		sum += intLine[len(intLine)-1]
	}
	return sum
}

func allZeros(intLine []int) bool {
	for _, v := range intLine {
		if v != 0 {
			return false
		}
	}
	return true
}

func diff(v []int) (ret []int, allZeros bool) {
	if len(v) <= 1 {
		panic("too short")
	}
	allZeros = true
	ret = make([]int, len(v)-1)
	for i := 0; i < len(v)-1; i++ {
		ret[i] = v[i+1] - v[i]
		if allZeros && ret[i] != 0 {
			allZeros = false
		}
	}
	return ret, allZeros
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
