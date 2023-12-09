package main

import (
	"aoc2023/common"
	"aoc2023/day09"
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

	// // day 6
	// lines := common.GetLinesFromFile("resources/06_race.txt", true, true)
	// // prod := day06.Part1(lines)
	// // fmt.Println(prod)
	// prod := day06.Part2(lines)
	// fmt.Println(prod)

	// // day 7
	// lines := common.GetLinesFromFile("resources/07_camel.txt", true, true)
	// // part1 := day07.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day07.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 8
	// lines := common.GetLinesFromFile("resources/08_leftright.txt", true, true)
	// // part1 := day08.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day08.Part2(lines)
	// fmt.Println("part2:", part2)

	// day 9
	lines := common.GetLinesFromFile("resources/09.txt", true, true)
	// part1 := day09.Part1(lines)
	// fmt.Println("part1:", part1)
	part2 := day09.Part2(lines)
	fmt.Println("part2:", part2)
}
