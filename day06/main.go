package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

type RecordRace struct {
	time     int
	distance int
}

func main() {
	input := util.ReadDay(6)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	races := parseInput(input)

	product := 1

	for _, race := range races {
		count := 0
		for t := 0; t <= race.time; t++ {
			if distance(t, race.time) > race.distance {
				count++
			}
		}
		product *= count
	}

	return product
}

func part2(input []string) int {
	races := parseInput([]string{
		strings.ReplaceAll(input[0], " ", ""),
		strings.ReplaceAll(input[1], " ", ""),
	})
	race := races[0]

	count := 0
	for t := 0; t <= race.time; t++ {
		if distance(t, race.time) > race.distance {
			count++
		}
	}

	return count
}

func parseInput(input []string) []RecordRace {
	results := []RecordRace{}

	timesStr, _ := strings.CutPrefix(input[0], "Time:")
	distancesStr, _ := strings.CutPrefix(input[1], "Distance:")
	times := parseNumbers(timesStr)
	distances := parseNumbers(distancesStr)

	for i := 0; i < len(times); i++ {
		results = append(results, RecordRace{times[i], distances[i]})
	}

	return results
}

func parseNumbers(list string) []int {
	result := []int{}

	for _, n := range strings.Split(list, " ") {
		if n == "" {
			continue
		}
		v, _ := strconv.Atoi(n)
		result = append(result, v)
	}

	return result
}

func distance(hold, raceTime int) int {
	return hold * (raceTime - hold)
}
