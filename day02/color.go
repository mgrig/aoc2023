package day02

import (
	"regexp"
	"strconv"
)

var reRed, _ = regexp.Compile(`.*?(\d+) red.*`)
var reGreen, _ = regexp.Compile(`.*?(\d+) green.*`)
var reBlue, _ = regexp.Compile(`.*?(\d+) blue.*`)

type Color struct {
	red, green, blue int
}

func NewColor(red, green, blue int) Color {
	return Color{
		red:   red,
		green: green,
		blue:  blue,
	}
}

// "3 blue, 4 red" > Color
func ParseColor(line string) Color {
	return NewColor(getValue(reRed, line), getValue(reGreen, line), getValue(reBlue, line))
}

func getValue(re *regexp.Regexp, line string) int {
	tokens := re.FindStringSubmatch(line)
	// fmt.Println(len(tokens), tokens)
	if len(tokens) == 2 {
		// found
		ret, _ := strconv.Atoi(tokens[1])
		return ret
	}
	return 0
}

func (c Color) smallerEqualThan(other Color) bool {
	return c.red <= other.red && c.green <= other.green && c.blue <= other.blue
}

func (c Color) Power() int {
	return c.red * c.green * c.blue
}
