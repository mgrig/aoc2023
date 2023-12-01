package day01

import (
	"aoc2023/common"
	"fmt"
	"regexp"
	"strconv"
)

func Part1(lines []string) int {
	reFirst, _ := regexp.Compile(`.*?(\d).*`) // lazy first .*, to get first digit
	reLast, _ := regexp.Compile(`.*(\d).*?`)  // lazy last .*, to get last digit
	sum := 0

	for _, line := range lines {
		first := common.GetOneRegexGroup(reFirst, line)
		last := common.GetOneRegexGroup(reLast, line)
		nr, _ := strconv.Atoi(first + last)
		sum += nr
	}

	return sum
}

var funny2proper = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part2(lines []string) int {
	reFirst, _ := regexp.Compile(`.*?(one|two|three|four|five|six|seven|eight|nine|\d).*`) // lazy first .*
	reLast, _ := regexp.Compile(`.*(one|two|three|four|five|six|seven|eight|nine|\d).*?`)  // lazy last .*
	sum := 0

	for _, line := range lines {
		first := stringFunnyToProper(common.GetOneRegexGroup(reFirst, line))
		last := stringFunnyToProper(common.GetOneRegexGroup(reLast, line))
		fmt.Println(first, last)
		nr, _ := strconv.Atoi(first + last)
		sum += nr
	}

	return sum
}

// "two" -> "2"
func stringFunnyToProper(str string) string {
	proper, exists := funny2proper[str]
	if exists {
		return proper
	}
	return str
}
