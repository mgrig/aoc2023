package day05

import (
	"aoc2023/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Part1(lines []string) int {
	reSeeds := regexp.MustCompile(`seeds: (.*)`)
	reMapHeader := regexp.MustCompile(`(.*?)-to-(.*?) map:`)

	var seeds []int
	allMaps := make(map[string](map[string]*Mapping))
	var currentMapping *Mapping = nil

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seedTokens := reSeeds.FindStringSubmatch(line)
			seeds = stringToIntArray(seedTokens[1])
			fmt.Println(seeds)
			continue
		}

		tokens := reMapHeader.FindStringSubmatch(line)
		if len(tokens) == 3 {
			currentMapping = NewMapping()

			destMap, exists := allMaps[tokens[1]]
			if !exists {
				allMaps[tokens[1]] = make(map[string]*Mapping)
				destMap = allMaps[tokens[1]]
			}

			destMap[tokens[2]] = currentMapping // should not already exist, but not checking atm
		} else {
			// normal line, add to existing mapping
			tokens := stringToIntArray(line)
			// fmt.Println(line, currentMapping)
			currentMapping.AddRange(*NewRangeMap(tokens[1], tokens[0], tokens[2]))
		}
	}
	// fmt.Println(allMaps)

	min := math.MaxInt
	for _, seed := range seeds {
		soil := allMaps["seed"]["soil"].Get(seed)
		fertilizer := allMaps["soil"]["fertilizer"].Get(soil)
		water := allMaps["fertilizer"]["water"].Get(fertilizer)
		light := allMaps["water"]["light"].Get(water)
		temp := allMaps["light"]["temperature"].Get(light)
		hum := allMaps["temperature"]["humidity"].Get(temp)
		location := allMaps["humidity"]["location"].Get(hum)

		if location < min {
			min = location
		}
	}

	return min
}

func Part2(lines []string) int {
	reSeeds := regexp.MustCompile(`seeds: (.*)`)
	reMapHeader := regexp.MustCompile(`(.*?)-to-(.*?) map:`)

	var seedsRaw []int
	allMaps := make(map[string](map[string]*Mapping))
	var currentMapping *Mapping = nil

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seedTokens := reSeeds.FindStringSubmatch(line)
			seedsRaw = stringToIntArray(seedTokens[1])
			continue
		}

		tokens := reMapHeader.FindStringSubmatch(line)
		if len(tokens) == 3 {
			currentMapping = NewMapping()

			destMap, exists := allMaps[tokens[1]]
			if !exists {
				allMaps[tokens[1]] = make(map[string]*Mapping)
				destMap = allMaps[tokens[1]]
			}

			destMap[tokens[2]] = currentMapping // should not already exist, but not checking atm
		} else {
			// normal line, add to existing mapping
			tokens := stringToIntArray(line)
			// fmt.Println(line, currentMapping)
			currentMapping.AddRange(*NewRangeMap(tokens[1], tokens[0], tokens[2]))
		}
	}
	// fmt.Println(allMaps)

	min := math.MaxInt
	for i := 0; i < len(seedsRaw); i += 2 {
		fmt.Println(seedsRaw[i], seedsRaw[i+1])
		for j := 0; j < seedsRaw[i+1]; j++ {
			seed := seedsRaw[i] + j
			soil := allMaps["seed"]["soil"].Get(seed)
			fertilizer := allMaps["soil"]["fertilizer"].Get(soil)
			water := allMaps["fertilizer"]["water"].Get(fertilizer)
			light := allMaps["water"]["light"].Get(water)
			temp := allMaps["light"]["temperature"].Get(light)
			hum := allMaps["temperature"]["humidity"].Get(temp)
			location := allMaps["humidity"]["location"].Get(hum)

			if location < min {
				min = location
			}
		}
	}

	return min
}

func stringToIntArray(line string) []int {
	tokens := strings.Split(strings.Trim(line, " "), " ")
	ret := make([]int, len(tokens))
	for i, token := range tokens {
		ret[i] = common.StringToInt(token)
	}
	return ret
}
