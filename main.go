package main

import (
	"aoc2023/common"
	"aoc2023/day06"
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

	// // day 4
	// lines := common.GetLinesFromFile("resources/04_scratch.txt", true, true)
	// // sum := day04.Part1(lines)
	// // fmt.Println("sum:", sum)
	// sum := day04.Part2(lines)
	// fmt.Println("sum2:", sum)

	// // day 5
	// lines := common.GetLinesFromFile("resources/05_garden.txt", true, true)
	// // min := day05.Part1(lines)
	// // fmt.Println("min:", min)
	// min := day05.Part2(lines)
	// fmt.Println("min:", min)

	// day 6
	lines := common.GetLinesFromFile("resources/06_race.txt", true, true)
	// prod := day06.Part1(lines)
	// fmt.Println(prod)
	prod := day06.Part2(lines)
	fmt.Println(prod)
}
