package main

import (
	"aoc2023/common"
	"aoc2023/day01"
	"fmt"
)

func main() {
	// day 1
	lines := common.GetLinesFromFile("resources/01_trebuchet.txt", true, true)
	// sum := day01.Part1(lines)
	// fmt.Println("sum:", sum)
	sum2 := day01.Part2(lines)
	fmt.Println("sum2:", sum2)
}
