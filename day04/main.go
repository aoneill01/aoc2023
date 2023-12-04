package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winning Numbers
	have    Numbers
}

type Numbers []int

func main() {
	input := util.ReadDay(4)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseInput(input []string) []Card {
	results := []Card{}

	for _, line := range input {
		results = append(results, parseLine(line))
	}

	return results
}

func parseLine(line string) Card {
	re := regexp.MustCompile(`^Card\s+(\d+):((\s+\d+)+) \|((\s+\d+)+)$`)
	matches := re.FindStringSubmatch(line)
	id, _ := strconv.Atoi(matches[1])

	return Card{
		id:      id,
		winning: parseNumbers(matches[2]),
		have:    parseNumbers(matches[4]),
	}
}

func parseNumbers(list string) Numbers {
	result := Numbers{}

	for _, n := range strings.Split(list, " ") {
		if n == "" {
			continue
		}
		v, _ := strconv.Atoi(n)
		result = append(result, v)
	}

	return result
}

func (n Numbers) contains(v int) bool {
	for _, current := range n {
		if current == v {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	cards := parseInput(input)
	result := 0

	for _, card := range cards {
		value := 0
		for _, w := range card.winning {
			if card.have.contains(w) {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}
		result += value
	}

	return result
}

func (c Card) winCount() int {
	result := 0

	for _, w := range c.winning {
		if c.have.contains(w) {
			result += 1
		}
	}

	return result
}

func part2(input []string) int {
	cards := parseInput(input)
	cardCounts := make([]int, len(cards))
	result := 0

	// Start with one of each card
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i, card := range cards {
		result += cardCounts[i]

		winCount := card.winCount()
		for j := 0; j < winCount; j++ {
			cardCounts[i+j+1] += cardCounts[i]
		}

	}

	return result
}
