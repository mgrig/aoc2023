package main

import (
	"aoc2023/common"
	"aoc2023/day04"
	"fmt"
)

func main() {
	// // day 1
	// lines := common.GetLinesFromFile("resources/01_trebuchet.txt", true, true)
	// // sum := day01.Part1(lines)
	// // fmt.Println("sum:", sum)
	// sum2 := day01.Part2(lines)
	// fmt.Println("sum2:", sum2)

	// // day 2
	// lines := common.GetLinesFromFile("resources/02_cubes.txt", true, true)
	// // sum := day02.Part1(lines)
	// // fmt.Println("sum:", sum)
	// sum := day02.Part2(lines)
	// fmt.Println("sum:", sum)

	// // day 3
	// lines := common.GetLinesFromFile("resources/03_engine.txt", true, true)
	// // sum := day03.Part1(lines)
	// // fmt.Println("sum:", sum)
	// sum := day03.Part2(lines)
	// fmt.Println("sum2:", sum)

	// day 4
	lines := common.GetLinesFromFile("resources/04_scratch.txt", true, true)
	sum := day04.Part1(lines)
	fmt.Println("sum:", sum)
}
