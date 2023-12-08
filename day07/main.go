package main

import (
	"aoc2023/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Card int
type Hand struct {
	cards            []Card
	bid              int
	strength         int
	wildcardStrength int
}

func main() {
	input := util.ReadDay(7)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	hands := parseInput(input)

	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]

		if hand1.strength > hand2.strength {
			return false
		}

		if hand2.strength > hand1.strength {
			return true
		}

		for i, card1 := range hand1.cards {
			card2 := hand2.cards[i]
			if card1 > card2 {
				return false
			}
			if card2 > card1 {
				return true
			}
		}

		return true
	})

	result := 0

	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}

	return result
}

func part2(input []string) int {
	hands := parseInput(input)

	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]

		if hand1.wildcardStrength > hand2.wildcardStrength {
			return false
		}

		if hand2.wildcardStrength > hand1.wildcardStrength {
			return true
		}

		for i, card1 := range hand1.cards {
			card2 := hand2.cards[i]
			if card1 == 11 {
				card1 = 1
			}
			if card2 == 11 {
				card2 = 1
			}
			if card1 > card2 {
				return false
			}
			if card2 > card1 {
				return true
			}
		}

		return true
	})

	result := 0

	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}

	return result
}

func parseCard(c string) Card {
	switch c {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		v, _ := strconv.Atoi(c)
		return Card(v)
	}
}

func parseLine(line string) Hand {
	hand := Hand{}
	parts := strings.Split(line, " ")
	cardStrings := strings.Split(parts[0], "")
	for _, cardString := range cardStrings {
		hand.cards = append(hand.cards, parseCard(cardString))
	}
	bid, _ := strconv.Atoi(parts[1])
	hand.bid = bid

	setStandardRank(&hand)
	setWildcardRank(&hand)

	return hand
}

func setStandardRank(hand *Hand) {
	countMap := map[Card]int{}
	for _, card := range hand.cards {
		countMap[card]++
	}
	counts := []int{}
	for _, v := range countMap {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if counts[0] == 5 {
		hand.strength = 7
	} else if counts[0] == 4 {
		hand.strength = 6
	} else if counts[0] == 3 {
		if counts[1] == 2 {
			hand.strength = 5
		} else {
			hand.strength = 4
		}
	} else if counts[0] == 2 {
		if counts[1] == 2 {
			hand.strength = 3
		} else {
			hand.strength = 2
		}
	} else {
		hand.strength = 1
	}
}

func setWildcardRank(hand *Hand) {
	countMap := map[Card]int{}
	var wildcardCount int
	for _, card := range hand.cards {
		if card == 11 {
			wildcardCount++
			continue
		}
		countMap[card]++
	}
	counts := []int{}
	for _, v := range countMap {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if len(counts) == 0 {
		counts = []int{0}
	}
	counts[0] += wildcardCount
	if counts[0] == 5 {
		hand.wildcardStrength = 7
	} else if counts[0] == 4 {
		hand.wildcardStrength = 6
	} else if counts[0] == 3 {
		if counts[1] == 2 {
			hand.wildcardStrength = 5
		} else {
			hand.wildcardStrength = 4
		}
	} else if counts[0] == 2 {
		if counts[1] == 2 {
			hand.wildcardStrength = 3
		} else {
			hand.wildcardStrength = 2
		}
	} else {
		hand.wildcardStrength = 1
	}
}

func parseInput(input []string) []Hand {
	hands := []Hand{}
	for _, line := range input {
		hands = append(hands, parseLine(line))
	}
	return hands
}
