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
	allMaps := make(map[string](map[string]*Fn))
	var currentFn *Fn = nil

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seedTokens := reSeeds.FindStringSubmatch(line)
			seedsRaw = stringToIntArray(seedTokens[1])
			continue
		}

		tokens := reMapHeader.FindStringSubmatch(line)
		if len(tokens) == 3 {
			currentFn = NewFn([]int{}, []int{})

			destMap, exists := allMaps[tokens[1]]
			if !exists {
				allMaps[tokens[1]] = make(map[string]*Fn)
				destMap = allMaps[tokens[1]]
			}

			destMap[tokens[2]] = currentFn // should not already exist, but not checking atm
		} else {
			// normal line, add to existing mapping
			tokens := stringToIntArray(line)
			// fmt.Println(line, currentMapping)
			currentFn.AddPointFromRaw(tokens[0], tokens[1], tokens[2])
		}
	}
	// fmt.Println(allMaps)

	seed2soil := allMaps["seed"]["soil"]
	soil2fert := allMaps["soil"]["fertilizer"]
	fert2water := allMaps["fertilizer"]["water"]
	water2light := allMaps["water"]["light"]
	light2temp := allMaps["light"]["temperature"]
	temp2hum := allMaps["temperature"]["humidity"]
	hum2location := allMaps["humidity"]["location"]

	// The magic is in the Compose function, which takes
	// y = f(x) and z = g(y) and creates a new interval-based function z = gf(x) = g(f(x))
	// Apply this repeatedly to get an interval based seed2location function.
	seed2location := Compose(*seed2soil, Compose(*soil2fert, Compose(*fert2water, Compose(*water2light, Compose(*light2temp, Compose(*temp2hum, *hum2location))))))

	// And then take the given seed ranges and consider them yet another function
	// (unity, aka f(x) = x), where we are only interested in the computation of
	// the combined domain ranges.
	range2seed := NewFn([]int{}, []int{})
	var dis DomainIntervals = make([]DomainInterval, 0)
	for i := 0; i < len(seedsRaw); i += 2 {
		fmt.Println(seedsRaw[i], seedsRaw[i+1])
		range2seed.AddPointFromRaw(seedsRaw[i], seedsRaw[i], seedsRaw[i+1])
		dis = append(dis, *NewDomainInterval(seedsRaw[i], seedsRaw[i]+seedsRaw[i+1]-1))
	}
	range2location := Compose(*range2seed, seed2location)

	// Each resulting domain range leads to the same(!) location value,
	// so it's enough to compute the min value once for each interval.
	// This is where the main time gain comes from.
	min := math.MaxInt
	for k, v := range range2location.offsets {
		if !dis.Contains(k) {
			continue
		}

		location := k + v

		if location == 0 {
			lower := seed2soil.domainLowerEqual(k)
			fmt.Println(k, v, lower)

		}

		if location < min {
			min = location
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
