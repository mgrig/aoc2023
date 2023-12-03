package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reGameSplit, _ = regexp.Compile(`Game (\d+): (.*)`)
var refColor = NewColor(12, 13, 14)

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		tokens := reGameSplit.FindStringSubmatch(line)
		if len(tokens) != 3 {
			panic("wrong game split")
		}

		game, _ := strconv.Atoi(tokens[1])
		gameColors := allGameColors(tokens[2])
		fmt.Println(game, gameColors)

		if isGamePossible(gameColors) {
			fmt.Println(game, "possible")
			sum += game
		}
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		tokens := reGameSplit.FindStringSubmatch(line)
		if len(tokens) != 3 {
			panic("wrong game split")
		}

		// game, _ := strconv.Atoi(tokens[1])
		gameColors := allGameColors(tokens[2])

		minColor := getMinColor(gameColors)
		// fmt.Println(game, minColor.Power())
		sum += minColor.Power()
	}

	return sum
}

func getMinColor(colors []Color) Color {
	ret := &Color{}

	for _, c := range colors {
		if c.red > ret.red {
			ret.red = c.red
		}
		if c.green > ret.green {
			ret.green = c.green
		}
		if c.blue > ret.blue {
			ret.blue = c.blue
		}
	}

	return *ret
}

func isGamePossible(gameColors []Color) bool {
	for _, gameColor := range gameColors {
		if !gameColor.smallerEqualThan(refColor) {
			return false
		}
	}
	return true
}

func allGameColors(line string) []Color {
	ret := make([]Color, 0)
	colorTokens := strings.Split(line, ";")
	for _, colorToken := range colorTokens {
		color := ParseColor(strings.Trim(colorToken, " "))
		ret = append(ret, color)
	}
	return ret
}
