package main

import (
	"aoc2023/common"
	"aoc2023/day21"
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

	// // day 9
	// lines := common.GetLinesFromFile("resources/09.txt", true, true)
	// // part1 := day09.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day09.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 10
	// lines := common.GetLinesFromFile("resources/10_loop.txt", true, true)
	// // part1 := day10.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day10.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 11
	// lines := common.GetLinesFromFile("resources/11.txt", true, true)
	// part1 := day11.Part1(lines, 2)
	// fmt.Println("part1:", part1)
	// part2 := day11.Part1(lines, 1000000)
	// fmt.Println("part2:", part2)

	// // day 12
	// lines := common.GetLinesFromFile("resources/12.txt", true, true)
	// // part1 := day12.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day12.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 13
	// lines := common.GetLinesFromFile("resources/13.txt", false, true)
	// part1 := day13.Part(lines, 0)
	// fmt.Println("part1:", part1)
	// part2 := day13.Part(lines, 1)
	// fmt.Println("part2:", part2)

	// // day 14
	// lines := common.GetLinesFromFile("resources/14.txt", true, true)
	// part1 := day14.Part1(lines)
	// fmt.Println("part1:", part1)
	// part2 := day14.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 15
	// lines := common.GetLinesFromFile("resources/15_hash.txt", true, true)
	// part1 := day15.Part1(lines[0])
	// fmt.Println("part1:", part1)
	// part2 := day15.Part2(lines[0])
	// fmt.Println("part2:", part2)

	// // day 16
	// lines := common.GetLinesFromFile("resources/16.txt", true, true)
	// // part1 := day16.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day16.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 17
	// lines := common.GetLinesFromFile("resources/17.txt", true, true)
	// // part1 := day17.Part1(lines)
	// // fmt.Println("part1:", part1)
	// part2 := day17.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 18
	// lines := common.GetLinesFromFile("resources/18_test.txt", true, true)
	// part1 := day18.Part1(lines)
	// fmt.Println("part1:", part1)
	// // Part 2 in Matlab (which solves both parts a lot faster)

	// // day 19
	// lines := common.GetLinesFromFile("resources/19.txt", true, true)
	// part1 := day19.Part1(lines)
	// fmt.Println("part 1:", part1)
	// part2 := day19.Part2(lines)
	// fmt.Println("part 2:", part2)

	// // day 20
	// lines := common.GetLinesFromFile("resources/20_cluster_fd_fn.txt", true, true)
	// // part1 := day20.Part1(lines)
	// // fmt.Println("part 1:", part1)
	// part2 := day20.Part2(lines)
	// fmt.Println("part 2:", part2)

	// day 21
	lines := common.GetLinesFromFile("resources/21.txt", true, true)
	part1 := day21.Part1(lines)
	fmt.Println("part 1:", part1)
	// part2 := day21.Part2(lines)
	// fmt.Println("part 2:", part2)
}
