package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type step struct {
	label       string
	operation   string
	focalLength int
}

type box []lens

type lens struct {
	label       string
	focalLength int
}

func (b box) process(s step) box {
	if s.operation == "-" {
		for i, l := range b {
			if l.label == s.label {
				for j := i; j < len(b)-1; j++ {
					b[j] = b[j+1]
				}
				return b[:len(b)-1]
			}
		}
		return b
	} else {
		for i, l := range b {
			if l.label == s.label {
				b[i] = lens{s.label, s.focalLength}
				return b
			}
		}
		return append(b, lens{s.label, s.focalLength})
	}
}

func main() {
	input := util.ReadDay(15)

	fmt.Println(part1(input[0]))
	fmt.Println(part2(input[0]))
}

func part1(input string) int {
	result := 0
	initSeq := strings.Split(input, ",")

	for _, step := range initSeq {
		result += hash(step)
	}

	return result
}

func part2(input string) int {
	result := 0
	initSeq := strings.Split(input, ",")
	var boxes [256]box

	for _, step := range initSeq {
		s := parseStep(step)
		boxId := hash(s.label)
		boxes[boxId] = boxes[boxId].process(s)
	}

	for i, b := range boxes {
		for j, l := range b {
			result += (i + 1) * (j + 1) * l.focalLength
		}
	}

	return result
}

func hash(val string) int {
	value := 0

	for _, b := range []byte(val) {
		value += int(b)
		value *= 17
		value = value % 256
	}

	return value
}

func parseStep(s string) step {
	re := regexp.MustCompile(`^([a-z]+)(=|-)(\d*)$`)
	matches := re.FindStringSubmatch(s)
	label := matches[1]
	operation := matches[2]
	var focalLength int
	if operation == "=" {
		focalLength, _ = strconv.Atoi(matches[3])
	}

	return step{label, operation, focalLength}
}
